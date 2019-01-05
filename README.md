# How to make a callback into C from Go

## Example: Count() method with a callback from Go
### C Code
Here's the `main.c`:
```c
#include "libcounter.h"
#include <stdio.h>

void onCount(int n) {
    printf("Count: %d\n", n);
}

int main(void) {
    int rc = Count(3, onCount); // implemented in Go
    printf("Done!\n");
    return rc;
}
```

The `Count()` method is implemented in Go, which is compiled as a shared library. The
method takes a number and a callback function. It calls the callback function, specified
number of times, with interval of 1-second, then exits.

### Go Code
Here's the implemention of `Count()` method in Go:

counter_api.go:
```go
package main

// typedef void (*count_cb)(int);
// void makeCallback(int sum, count_cb cb);
import "C"

//export Count
func Count(n C.int, cb C.count_cb) int {
	goCb := func(n int) {
		C.makeCallback(C.int(n), cb)
	}
	return count(int(n), goCb)
}
```

counter.go:
```go
package main

// typedef void (*count_cb)(int);
// extern void bridge(int sum, count_cb cb) {
//     cb(sum);
// }
import "C"
import (
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

func main() {}
```
## How to Build
First build a shared library from Go code.
```sh
go build -v -buildmode=c-shared -o libcounter.so *.go
```

Then build a C executable using the library.
```sh
gcc -o main main.c libcounter.so
```

> Makefile is provided. For a quick build, type `make`.

To run:
```
$ ./main
Count: 1
Count: 2
Count: 3
Done!
```

Also try:
```
$ go run counter.go
```
It demonstrate that by putting C-typed (binding) code as a separate file
(seen as counter_api.go), the rest of `pure` Go code become testable and
runnable within Go land.

## Comments / Tips
* No blank-line between C-code comment lines and `import "C"`
* Go code can call C-function in the comment lines. But;
* Go code can NOT directly call function *pointer*.
* Go method meant to be called by C must have `//export Xxx` right above the function.
* The `*.go` file that contains `//export Xxx` cannot have definition in the comment line, declaration only. (see counter.go)
* For that reason, use a sperate go file to place the definition. (see counter.go)
* No space between `//` and `expose`. (`// export Xxx` is tempting, but wound't work)
* For more details, see [The Go Programming Language, Command cgo](https://golang.org/cmd/cgo/).


