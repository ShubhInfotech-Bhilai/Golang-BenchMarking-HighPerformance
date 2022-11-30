package main
import (
	"runtime/trace"
	"os"
  )
  //go tool trace binary trace.out
  func main(){
 f,_:= os.Create("trace.out")
  defer f.Close()
  _= trace.Start(f)
  defer trace.Stop()

  }
  