package main

// typedef void (*count_cb)(int);
// void bridge(int sum, count_cb cb);
import "C"

// Count adds two numbers
// IMPORTANT: The following "export Count" mark this method as
// a function exposed to C, and will be put in the auto-generated
// header file.
//export Count
func Count(n C.int, cb C.count_cb) int {
	goCb := func(n int) {
		C.bridge(C.int(n), cb)
	}
	return count(int(n), goCb)
}
