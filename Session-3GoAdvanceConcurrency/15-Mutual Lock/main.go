package main
import (
	"fmt"
	"context"
	//"os"
	//"strconv"
	"time"
	"golang.org/x/sync/semaphore"
)
var Worker=4
var sem=semaphore.NewWeighted(int64(Worker))
var njobs=20

func worker(n int )int {
	sq:=n*n
	time.Sleep(time.Second)
	return sq
}

func main(){
	var result =make([]int , njobs)
	 var err error
	//TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear 
	//which Context to use or it is not yet available
	// (because the surrounding function has not yet been extended to accept a Context parameter).
	 ctx:=context.TODO()
	  for i:= range result{
		err =sem.Acquire(ctx,1)
		if err!=nil{
			fmt.Println("Sephore Error")
			break
		}
		go func(i int ){
			defer sem.Release(1)
			temp:=worker(i)
			result[i]=temp

		}(i)
	  }
	  err =sem.Acquire(ctx,int64(Worker))
	  if err !=nil{
		fmt.Println("Sephore Error"  ,err)
	  }
	  for k,v:= range result{
		fmt.Println(k , "->" , v)
	  }




}