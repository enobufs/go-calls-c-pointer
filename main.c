#include "libcounter.h"
#include <stdio.h>


void onCount(int n) {
    printf("Count: %d\n", n);
}

int main(void) {
    int rc = Count(3, onCount);
    printf("Done!\n");
    return rc;
}

