package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
    var m runtime.MemStats
	debug.SetGCPercent(-1)
	a := make([]*int, 1e2)
	
	runtime.ReadMemStats(&m)
	fmt.Println(a[0])
	fmt.Println("Memory before GC ", m.HeapInuse)
	time.Sleep(5*time.Second)
	debug.SetGCPercent(100)
	for i:=0;i<=100;i++{
		runtime.GC()
	}
	runtime.ReadMemStats(&m)
	fmt.Println("Memory After GC " ,m.HeapInuse)

	/*
	for i := 0; i <= 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Println("GC took %s\n", time.Since(start))
	}
	runtime.KeepAlive(a)
	*/

}
