// +build gperf

package cgogperf

/*
#cgo CXXFLAGS: -I ../include
#cgo CFLAGS: -I ../include
#cgo LDFLAGS: -L ../lib -ltcmalloc -lprofiler -lstdc++
#include <stdlib.h>
#include <gperftools/profiler.h>
#include <gperftools/heap-profiler.h>

#include "heap_checker.h"
*/
import "C"
import (
	"log"
	"unsafe"

	gogperf "github.com/amazingchow/go-gperftools"
)

func init() {
	gogperf.CGOHeapProfiler = &cgoHeapProfiler{}
	gogperf.CGOHeapLeakChecker = &cgoHeapLeakChecker{}
}

//********** HeapProfiler **********//
type cgoHeapProfiler struct {
	hasStarted bool
}

func (c *cgoHeapProfiler) Start(name string) {
	if c.Started() {
		log.Fatal("CGOHeapProfiler has started already!!!")
	}
	startHeapProfiler(name)
	c.hasStarted = true
}

func (c *cgoHeapProfiler) Started() bool {
	return c.hasStarted
}

func (c *cgoHeapProfiler) Stop(name string) {
	if !c.Started() {
		log.Fatal("CGOHeapProfiler has not started yet!!!")
	}
	dumpHeapProfiler(name)
	stopHeapProfiler()
	c.hasStarted = false
}

func startHeapProfiler(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.HeapProfilerStart(cName)
}

func dumpHeapProfiler(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.HeapProfilerDump(cName)
}

func stopHeapProfiler() {
	C.HeapProfilerStop()
}

//********** HeapLeakChecker **********//
type cgoHeapLeakChecker struct {
	handler C.HeapLeakCheckerCGO
}

func (c *cgoHeapLeakChecker) Start(name string) {
	if c.Started() {
		log.Fatal("cgoHeapLeakChecker has started already!!!")
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	c.handler = C.start(cName)
}

func (c *cgoHeapLeakChecker) Started() bool {
	return c.handler != nil
}

func (c *cgoHeapLeakChecker) Stop(_ string) {
	if !c.Started() {
		log.Fatal("cgoHeapLeakChecker has not started yet!!!")
	}
	C.stop(c.handler)
	c.handler = nil
}
