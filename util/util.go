// Author: sheppard(ysf1026@gmail.com) 2013-11-22

package util

import (
	"fmt"
	"os"
	"runtime"
)

func CheckFatal(err error) {
	if err != nil {
		var stack []byte = make([]byte, 2048)
		runtime.Stack(stack, false)
		fmt.Fprintf(os.Stderr, "Fatal: %s [stack]: %s", err.Error(), string(stack))
		os.Exit(2)
	}
}

func CheckError(err error) {
	if err != nil {
		var stack []byte = make([]byte, 2048)
		runtime.Stack(stack, false)
		fmt.Fprintf(os.Stderr, "Error: %s [stack]: %s", err.Error(), string(stack))
	}
}
