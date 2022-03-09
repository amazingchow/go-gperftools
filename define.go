package gogperf

import (
	"os"
)

var (
	// CGOHeapProfiler 自定义堆内存检测工具.
	CGOHeapProfiler GPerfHelper
	// EnableCGOHeapProfiler 是否检测堆内存使用情况.
	EnableCGOHeapProfiler bool
	// CGOHeapLeakChecker 自定义堆内存泄露检测工具.
	CGOHeapLeakChecker GPerfHelper
	// EnableCGOHeapLeakChecker 是否检测堆内存泄露情况.
	EnableCGOHeapLeakChecker bool
)

func init() {
	// HEAPPROFILE is the path to dist profile, eg: /tmp/mybin.hprof.
	// More info: https://gperftools.github.io/gperftools/heapprofile.html.
	if len(os.Getenv("HEAPPROFILE")) > 0 {
		EnableCGOHeapProfiler = true
	}

	// HEAPCHECK is the check level, eg: normal.
	// More info: https://gperftools.github.io/gperftools/heap_checker.html.
	if len(os.Getenv("HEAPCHECK")) > 0 {
		EnableCGOHeapLeakChecker = true
	}
}

// StartGPerfHelper 开启gperftools检测工具.
func StartGPerfHelper(name string) {
	if EnableCGOHeapProfiler {
		CGOHeapProfiler.Start(name)
	}
	if EnableCGOHeapLeakChecker {
		CGOHeapLeakChecker.Start(name)
	}
}

// StopGPerfHelper 关闭gperftools检测工具.
func StopGPerfHelper(name string) {
	if EnableCGOHeapProfiler {
		CGOHeapProfiler.Stop(name)
	}
	if EnableCGOHeapLeakChecker {
		CGOHeapLeakChecker.Stop(name)
	}
}
