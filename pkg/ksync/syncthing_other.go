// +build !linux

package ksync

import (
	"syscall"
)

var syncthingProcAttr = &syscall.SysProcAttr{}
