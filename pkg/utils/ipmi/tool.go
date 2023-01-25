// Copyright 2022 Listware

package ipmi

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"

	"git.fg-tech.ru/listware/inventory-app/pkg/utils/driver"
)

// https://webcache.googleusercontent.com/search?q=cache:ohTuxLGBYaEJ:https://computercheese.blogspot.com/2014/03/&cd=2&hl=ru&ct=clnk&gl=ru
var (
	// Tool \\
	Tool = &tool{}
	// Driver \\
	Driver = driver.Driver("ipmi_si")
)

type tool struct{}

func (t *tool) MC() MC {
	mc := &mgmtController{tool: t}
	return mc
}

func (t *tool) exec(method string, args ...string) (*bytes.Buffer, error) {
	a := []string{method}
	a = append(a, args...)
	c := exec.Command("ipmitool", a...)
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	c.Stdout = stdout
	c.Stderr = stderr
	if err := c.Run(); err != nil {
		return nil, fmt.Errorf("%s: %w", stderr.String(), err)
	}
	return stdout, nil
}

// MC - management controller
type MC interface {
	GUID() (string, error)
	IP() (string, error)
	Temp() ([4]int, error)
	Reset() error // rtype: <warm|cold>
}

type mgmtController struct {
	tool *tool
}

func (mc *mgmtController) GUIDOld() (s string, err error) {
	bf, err := mc.tool.exec("mc", "guid")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(bf)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "System GUID") {
			kv := strings.Split(txt, ": ")
			if len(kv) != 2 {
				err = fmt.Errorf("failed to parse ipmitool output: %s", txt)
				return
			}
			s = kv[1]
			return
		}
	}
	err = fmt.Errorf("failed to parse ipmitool output: %s", bf)
	return
}

func (mc *mgmtController) GUID() (s string, err error) {
	buf, err := mc.tool.exec("raw", "0x06", "0x037")
	if err != nil {
		return
	}
	rawString := normalize(buf.Bytes())
	output := make([]string, 5)
	output[0] = revert(rawString, 0, 7)
	output[1] = revert(rawString, 1+7, 7+4)
	output[2] = revert(rawString, 1+7+4, 7+4+4)
	output[3] = string(rawString[1+7+4+4 : 1+7+4+4+4])
	output[4] = string(rawString[1+7+4+4+4 : 1+7+4+4+4+12])
	return strings.Join(output, "-"), err
}

func revert(r []byte, f, t int) (result string) {
	for i := t; i >= f; i -= 2 {
		result += string(r[i-1]) + string(r[i])
	}
	return
}

func normalize(src []byte) []byte {
	str := strings.ReplaceAll(string(src), " ", "")
	return []byte(strings.ReplaceAll(str, "\n", ""))
}

func (mc *mgmtController) IP() (s string, err error) {
	buf, err := mc.tool.exec("raw", "0x0C", "0x02", "0x1", "0x3", "0x0", "0x0")
	if err != nil {
		return
	}
	rawString := normalize(buf.Bytes())
	dst := make([]byte, hex.DecodedLen(len(rawString)))
	if _, err = hex.Decode(dst, rawString); err != nil {
		return
	}

	if dst[0] != 0x11 {
		err = fmt.Errorf("bad code")
		return
	}
	s = fmt.Sprintf("%d.%d.%d.%d", dst[1], dst[2], dst[3], dst[4])
	return
}

func (mc *mgmtController) Reset() (err error) {
	_, err = mc.tool.exec("mc", "reset", "cold")
	return
}

// ipmitool -b 0x06 -t 0x2c raw 0x2e 0x4b 0x57 0x01 0x00 0x0f 0xff 0xff 0xff 0xff 0xff 0xff 0xff 0xff
// -b channel
// -t address
// header: 57 01 00
// cpus (2 из 4 Шт): 2c 2f ff ff
// mem (посмотри в IpmiTemperatureStats::SystemTopology::setChannels):
//
//	2d ff ff ff 2d ff ff ff 2d
//	ff ff ff 2d ff ff ff 2d ff ff ff 2d ff ff ff 2e
//	ff ff ff 2d ff ff ff ff ff ff ff ff ff ff ff ff
//	ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff
//	ff ff ff ff ff ff ff
func (mc *mgmtController) Temp() (cpuTemp [4]int, err error) {
	buf, err := mc.tool.exec("-b", "0x06", "-t", "0x2c", "raw", "0x2e", "0x4b", "0x57", "0x01", "0x00", "0x0f", "0xff", "0xff", "0xff", "0xff", "0xff", "0xff", "0xff", "0xff")
	if err != nil {
		return
	}
	rawString := normalize(buf.Bytes())
	dst := make([]byte, hex.DecodedLen(len(rawString)))
	if _, err = hex.Decode(dst, rawString); err != nil {
		return
	}

	if dst[0] != 0x57 || dst[1] != 0x01 || dst[2] != 0x00 {
		err = fmt.Errorf("bad code")
		return
	}

	cpuTemp[0] = int(dst[3])
	cpuTemp[1] = int(dst[4])
	cpuTemp[2] = int(dst[5])
	cpuTemp[3] = int(dst[6])

	return
}
