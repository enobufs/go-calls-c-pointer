# How to make a callback into C from Go

## What does this example do?
Here's the `main.c`:
```c
#include "libcounter.h"
#include <stdio.h>

void callback_for_count(int n) {
    printf("Count: %d\n", n);
}

int main(void) {
    int rc = Count(3, callback_for_count); // implemented in Go
    printf("Done!\n");
    return rc;
}

```
The `Count()` method is implemented in Go, which is compiled as a shared library. The
method takes a number and a callback function. It calls the callback function, specified
number of times, with interval of 1-second, then exits.

## How to run
### Build
```
make
```

or

```
go build -v -buildmode=c-shared -o libcounter.so *.go
gcc -o main main.c libcounter.so
```

### Run
```
$ ./main
Count: 1
Count: 2
Count: 3
Done!
```

## Comments
* No space between C-code comment lines and `import "C"`
* Go code can call C-function in the comment lines. But;
* Go code can NOT directly call function *pointer*.
* Go method meant to be called by C must have `//export Xxx` right above the function.
* The `*.go` file that contains `//export Xxx` cannot have definition in the comment line, declaration only. (see counter.go)
* No space between `//` and `expose`. (`// export Xxx` is tempting, but no good)
* For that reason, use a sperate go file to place the definition. (see counter.go)
* For more details, see [The Go Programming Language, Command cgo](https://golang.org/cmd/cgo/).


