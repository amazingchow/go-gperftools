# go-gperftools

## How-To-Use

```golang
// +build gperf

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	gogperf "github.com/amazingchow/go-gperftools"
)

func init() {
	go func() {
		if gogperf.EnableCGOHeapProfiler && gogperf.CGOHeapProfiler == nil {
			log.Fatal("`gperf` build flag not enabled")
		}
		if gogperf.EnableCGOHeapProfiler && gogperf.CGOHeapProfiler.Started() {
			gogperf.CGOHeapProfiler.Stop("gperf-heap-profiler")
		}
		if gogperf.EnableCGOHeapLeakChecker && gogperf.CGOHeapLeakChecker == nil {
			log.Fatal("`gperf` build flag not enabled")
		}
		if gogperf.EnableCGOHeapLeakChecker && gogperf.CGOHeapLeakChecker.Started() {
			gogperf.CGOHeapLeakChecker.Stop("gperf-heap-profiler")
		}

		http.HandleFunc("/cgo/heapprofiler/start", func(w http.ResponseWriter, r *http.Request) {
			if gogperf.CGOHeapProfiler != nil {
				if !gogperf.CGOHeapProfiler.Started() {
					log.Println("start gperf-heap-profiler")
					gogperf.CGOHeapProfiler.Start(fmt.Sprintf("%v", time.Now()))
				}
			}
		})
		http.HandleFunc("/cgo/heapprofiler/stop", func(w http.ResponseWriter, r *http.Request) {
			if gogperf.CGOHeapProfiler != nil {
				if gogperf.CGOHeapProfiler.Started() {
					log.Println("stop gperf-heap-profiler")
					gogperf.CGOHeapProfiler.Stop(fmt.Sprintf("%v", time.Now()))
				}
			}
		})
		http.HandleFunc("/cgo/heapleakchecker/start", func(w http.ResponseWriter, r *http.Request) {
			if gogperf.CGOHeapLeakChecker != nil {
				if !gogperf.CGOHeapLeakChecker.Started() {
					log.Println("start heap check metrics")
					gogperf.CGOHeapLeakChecker.Start("")
				}
			}
		})
		http.HandleFunc("/cgo/heapleakchecker/stop", func(w http.ResponseWriter, r *http.Request) {
			if gogperf.CGOHeapLeakChecker != nil {
				if !gogperf.CGOHeapLeakChecker.Started() {
					log.Println("start gperf-heap-leak-checker")
					gogperf.CGOHeapLeakChecker.Start("gperf-heap-leak-checker")
				} else {
					log.Println("stop gperf-heap-leak-checker")
					gogperf.CGOHeapLeakChecker.Stop("gperf-heap-leak-checker")
				}
			}
		})
		log.Fatal(http.ListenAndServe(":8082", nil))
	}()
}
```
