package main

import (
	"runtime"
	"time"
)

// #include <unistd.h>
// #include <errno.h>
// void setErrno() { errno = 1; }
import "C"

func main() {
	for i := 0; i < runtime.NumCPU()*2; i++ { // make sure it run goroutine on all CPUs
		go func() {
			time.Sleep(500 * time.Millisecond)
			// C.setErrno()
		}()
	}
	if _, err := C.sleep(C.uint(1)); err != nil {
		panic(err)
	}
}
