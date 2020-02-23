package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

// #include <unistd.h>
// #include <errno.h>
// void wait() { usleep(2000000); }
import "C"

func main() {
	for i := 0; i < runtime.NumCPU()*2; i++ { // run goroutine on all CPUs
		go func() {
			for {
				_, _ = net.Dial("tcp", "localhost:12345") // set errno
				time.Sleep(time.Millisecond)
			}
		}()
	}
	_, err := C.wait()
	fmt.Printf("CGO: %v\n", err)
	if err != nil {
		os.Exit(1)
	}
}
