package main

/*
#cgo LDFLAGS: -l winmm
#include <windows.h>
#include <stdio.h>

#pragma comment(lib, "Winmm.lib")

int MCIExecute(const char* cmd) {
	return mciSendString(cmd, "", 0, NULL);
}
*/
import "C"

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Execute(cmd string) int {
	code := C.MCIExecute(C.CString(cmd))
	return int(code)
}

func main() {
	startSecond, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	endSecond, err := strconv.Atoi(os.Args[3])

	Execute(fmt.Sprintf("open %s alias music", os.Args[1]))
	time.Sleep(time.Duration(startSecond) * time.Second)
	Execute("play music")
	time.Sleep(time.Duration(endSecond) * time.Second)
}
