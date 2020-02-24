package main

import (
	"runtime"
	"time"
)

// #include <unistd.h>
// #include <errno.h>
// #include <time.h>
// void setErrno() { errno = 1; }
// void wait() {
//   time_t n, s = time(0);
//   do {
//     time(&n);
//   } while(difftime(n, s) < 1.0);
// }
import "C"

func main() {
	for i := 0; i < runtime.NumCPU()*2; i++ { // make sure it run goroutine on all CPUs
		go func() {
			time.Sleep(500 * time.Millisecond)
			// C.setErrno()
		}()
	}
	if _, err := C.wait(); err != nil {
		panic(err)
	}
}
