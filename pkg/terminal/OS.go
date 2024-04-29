package terminal

import (
	"fmt"
	"runtime"
)

func GetOS() string {
	os := runtime.GOOS
	fmt.Println(os)
	return os

}
