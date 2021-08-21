package pprofx

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	sigar "github.com/cloudfoundry/gosigar"
	"github.com/peractio/gdk/pkg/converter"
	"github.com/peractio/gdk/pkg/jsonx"
)

const SleepDuration = time.Second * 10

func New(addresses ...string) {
	addr := ":6060"
	if len(addresses) > 0 {
		addr = addresses[0]
	}

	runtime.SetBlockProfileRate(1)

	// We will keep trying to run the server.
	// If the current address is busy,
	// sleep then try again until the address has become available.
	go func() {
		server := http.NewServeMux()
		server.HandleFunc("/", Summary)
		server.HandleFunc("/debug/pprof/", pprof.Index)
		server.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		server.HandleFunc("/debug/pprof/profile", pprof.Profile)
		server.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		server.HandleFunc("/debug/pprof/trace", pprof.Trace)

		for {
			if err := http.ListenAndServe(addr, server); err != nil {
				time.Sleep(SleepDuration)
			}
		}
	}()
}

func Summary(w http.ResponseWriter, _ *http.Request) {
	uptime := sigar.Uptime{}
	_ = uptime.Get()

	concreteSigar := sigar.ConcreteSigar{}
	avg, _ := concreteSigar.GetLoadAverage()
	mem, _ := concreteSigar.GetMem()
	swap, _ := concreteSigar.GetSwap()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	resp := map[string]interface{}{
		"htop": map[string]string{
			"uptime":      uptime.Format(),
			"load_avg_1":  fmt.Sprintf("%.2f", avg.One),
			"load_avg_5":  fmt.Sprintf("%.2f", avg.Five),
			"load_avg_15": fmt.Sprintf("%.2f", avg.Fifteen),
		},
		"free": map[string]interface{}{
			"mem": map[string]string{
				"total":           converter.ByteSize(mem.Total),
				"used":            converter.ByteSize(mem.ActualUsed),
				"used_percentage": converter.Percentage(mem.ActualUsed, mem.Total),
				"free_percentage": converter.Percentage(mem.ActualFree, mem.Total),
				"free":            converter.ByteSize(mem.ActualFree),
			},
			"swap": map[string]string{
				"total":           converter.ByteSize(swap.Total),
				"used":            converter.ByteSize(swap.Used),
				"used_percentage": converter.Percentage(swap.Used, swap.Total),
				"free_percentage": converter.Percentage(swap.Free, swap.Total),
				"free":            converter.ByteSize(swap.Free),
			},
		},
		"mem_stats": map[string]string{
			"alloc":       converter.ByteSize(m.Alloc),
			"total_alloc": converter.ByteSize(m.TotalAlloc),
			"sys":         converter.ByteSize(m.Sys),
			"num_gc":      converter.String(m.NumGC),
			"live_objs":   converter.String(m.Mallocs - m.Frees),
		},
		"references": map[string]string{
			"alloc":       "currently allocated number of bytes on the heap",
			"total_alloc": "cumulative max bytes allocated on the heap (will not decrease)",
			"sys":         "total memory obtained from the OS",
			"num_gc":      "number of completed GC cycles",
			"live_objs":   "number of allocations, deallocations, and live objects (mallocs - frees)",
		},
	}

	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = jsonx.NewEncoder(w).Encode(resp)
}
