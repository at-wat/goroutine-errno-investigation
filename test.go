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
	for i := 0; i < runtime.NumCPU()*2; i++ { // make sure it run goroutine on all CPUs
		go func() {
			time.Sleep(500 * time.Millisecond)
			_, _ = net.Dial("tcp", "localhost:12345") // dst port is not listened
		}()
	}
	_, err := C.usleep(C.uint(1000000))
	fmt.Printf("error from CGO: %v\n", err)
	if err != nil {
		os.Exit(1)
	}
}
