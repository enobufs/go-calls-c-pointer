package main

// typedef void (*count_cb)(int);
// void makeCallback(int sum, count_cb cb);
import "C"

// Count calls callback speicifed number of times, then exits.
// IMPORTANT: The following "export Count" mark this method as
// a function exposed to C, and will be put in the auto-generated
// header file.
//export Count
func Count(n C.int, cb C.count_cb) int {

	// With a help of closure, we can make counter.go a totally
	// pure Go code.
	// (No C-types used in the function signature for goCb.)
	goCb := func(n int) {
		// Go cannot call C-function pointers.. Instead, use
		// a C-function to have it call the function pointer.
		C.makeCallback(C.int(n), cb)
	}
	return count(int(n), goCb)
}
