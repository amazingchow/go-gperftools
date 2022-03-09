#include <stdio.h>
#include <stdlib.h>
#include <gperftools/heap-checker.h>

#include "heap_checker.h"

HeapLeakCheckerCGO start(char* name) {
    HeapLeakChecker *handler = new HeapLeakChecker(name);
    return handler;
}

void stop(HeapLeakCheckerCGO handler) {
    HeapLeakChecker *heap_leak_checker = (HeapLeakChecker *)handler;
    if (!heap_leak_checker->NoLeaks()) {
        fprintf(stderr, "heap memory is leaking...\n");
    }
	delete heap_leak_checker;
}
