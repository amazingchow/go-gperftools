#include <stdio.h>
#include <stdlib.h>
#include <gperftools/heap-checker.h>

#include "heap_checker.h"

HeapLeakCheckerCGO start(char* name) {
    HeapLeakChecker *heap_checker = new HeapLeakChecker(name);
    return heap_checker;
}

void stop(HeapLeakCheckerCGO i) {
    HeapLeakChecker *heapC = (HeapLeakChecker *)i;
    if (!heapC->NoLeaks()) {
        fprintf(stderr, "heap memory leak\n");
    }
	delete heapC;
}
