package main

import (
	"runtime"
)

// #include <unistd.h>
// #include <errno.h>
// #include <time.h>
// void setErrno() { while (1) { errno = 1; } }
// void wait() {
//   time_t n = time(0), s = time(0);
//   do {
//     time(&n);
//   } while(difftime(n, s) < 5.0);
// }
import "C"

func main() {
	for i := 0; i < runtime.NumCPU()*2; i++ { // make sure to run on all CPUs
		go func() {
			C.setErrno()
		}()
	}
	// sleep/usleep on always sets errno=60 on Darwin
	if _, err := C.wait(); err != nil {
		panic(err)
	}
}
