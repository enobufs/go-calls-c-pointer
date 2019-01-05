#include "libcounter.h"
#include <stdio.h>


void callback_for_count(int n) {
    printf("Count: %d\n", n);
}

int main(void) {
    int rc = Count(3, callback_for_count);
    printf("Done!\n");
    return rc;
}

