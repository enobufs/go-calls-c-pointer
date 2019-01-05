.PHONY: all clean

default: main.c libcounter.so
	gcc -o main $?

libcounter.so: counter_api.go counter.go
	go build -buildmode=c-shared -o $@ *.go

clean:
	rm -f main libcounter.so libcounter.h
