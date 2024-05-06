package terminal

import (
	"runtime"
)

func GetOS() string {
	os := runtime.GOOS
	return os

}
