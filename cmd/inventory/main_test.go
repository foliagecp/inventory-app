// Copyright 2023 NJWS Inc.

package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(t *testing.T) {
	defer goleak.VerifyNone(t)
}
