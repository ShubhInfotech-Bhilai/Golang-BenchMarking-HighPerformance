
package main
//GOOS=linux GOARCH=amd64 go tool compile -S direct_topfunc_call.go
//GOOS=windows GOARCH=amd64 go tool compile -S direct_topfunc_call.go
//go tool compile -S GoAssembly.go > assembly.asm
func add(a, b int32) (int32, bool) { return a + b, true }

func main() { add(10, 32) }