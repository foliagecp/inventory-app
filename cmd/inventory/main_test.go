// Copyright 2022 Listware

package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(t *testing.T) {
	defer goleak.VerifyNone(t)
}
