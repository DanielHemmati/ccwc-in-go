package main

import (
	"fmt"
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// a method for displaying memory usage
func Mem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v MiB | ", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc: %v MiB | ", bToMb(m.TotalAlloc))
	fmt.Printf("\tHeapAlloc: %v MiB | ", bToMb(m.HeapAlloc))
	fmt.Printf("\tHeapSys: %v MiB | ", bToMb(m.HeapSys))
	fmt.Printf("\tNumGC: %v\n", m.NumGC)
}
