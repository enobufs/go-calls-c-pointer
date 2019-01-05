package main

// typedef void (*count_cb)(int);
// extern void bridge(int sum, count_cb cb) {
//     cb(sum);
// }
import "C"

// Above lines have nothing to do with the code below.
// It just implements the C-declaration found in
// counter_api.go. It is because counter_api.go contains
// `//export Xxx` and it will be used to generate a C header
// file `libcounter.h`.

import (
	"fmt"
	"time"
)

type countCb func(n int)

func count(n int, goCb countCb) int {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second * 1)
		goCb(i + 1)
	}

	return 0
}

// This main function won't be called when compiled as a shared
// library. But you can use it to run a quick standalone test.
// For this reason, it is recommended that a separate file be
// used for exported methods with C-types.
func main() {
	fmt.Println("Running standalone..")
	count(3, func(n int) {
		fmt.Printf("count %d\n", n)
	})
}
