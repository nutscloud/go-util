package net

import (
	"runtime"
)

var Bridge Bridger

type Bridger interface {
	GetBridgeInterfaces(br string) ([]string, error)
	GetAllBridge() ([]string, error)
}

func init() {
	switch runtime.GOOS {
	case "linux":
		Bridge = &linuxBridge{}
	}
}
