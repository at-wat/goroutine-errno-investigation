package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

// #include <unistd.h>
import "C"

func main() {
	for i := 0; i < runtime.NumCPU()*2; i++ { // run goroutine on all CPUs
		go func() {
			for {
				_, _ = net.Dial("tcp", "localhost:12345") // dst port is not listened
				time.Sleep(time.Millisecond)
			}
		}()
	}
	_, err := C.usleep(C.uint(1000000))
	fmt.Printf("error from CGO: %v\n", err)
	if err != nil {
		os.Exit(1)
	}
}
