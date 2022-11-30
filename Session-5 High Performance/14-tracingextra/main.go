package main

import (
    "os"
    "runtime/trace"
)
//go tool trace trace.out
func main() {
	f, _ := os.Create("trace.out")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()
   
}