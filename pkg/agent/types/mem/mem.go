// Copyright 2022 Listware

package mem

import (
	"strings"
)

// Ram module profile
type MemoryDevice struct {
	Device     string `json:"device,omitempty"`
	Bank       string `json:"bank,omitempty"`
	FormFactor string `json:"form-factor,omitempty"`
	Type       string `json:"type,omitempty"`

	Size         uint16 `json:"size,omitempty"`
	Speed        uint16 `json:"speed,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Serial       string `json:"serial,omitempty"`
	Part         string `json:"part,omitempty"`
}

func (m *MemoryDevice) Name() string {
	return strings.ToLower(strings.ReplaceAll(m.Device, "_", "-"))
}
