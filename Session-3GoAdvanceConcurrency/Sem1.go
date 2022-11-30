package main

import (
	"context"
	"math/rand"
	"errors"
	"fmt"

	"time"

	"golang.org/x/sync/semaphore"
)
func main(){
	ctx,cancle:=context.WithCancel(context.Background())

	defer cancle()

	var (
		maxWorker= 1000
		routine= 100000
		sem= semaphore.NewWeighted(int64(maxWorker))
	)
	errs:=make(chan error,2)
    fmt.Println("Main Starting")
	for i:=0;i<=routine;i++{
		err :=sem.Acquire(ctx,1)
		if err!=nil{
			fmt.Println("Sephore Error", err)
			break
		}
          //----Go Routine Start--------------------------------
		  go func(n int ){
			defer sem.Release(1)
			
			select{
			case <-ctx.Done():
				return
			default:

			}
			/* ===================Go routine Called "  "  */ 
			err:= Routine(n)
			if err!=nil{
				errs<-fmt.Errorf("Worker %d ,error %v\n" , i,err )
				cancle()
				return
			}
		}(i)
	}	
        // Semaphore Reported
		if err:=sem.Acquire(ctx,int64(maxWorker));err!=nil{
			fmt.Println("Failed to Acquire Sephore Error", err)

		}
       // error reported by goroutine
		if ctx.Err()!=nil{
			fmt.Printf("Go routine Error %v ",<-errs )
		}else{
			fmt.Printf("Go routine ok %v ",<-errs )
		}

		fmt.Print("Exit from Main")

	}
	func Routine(n int )error{
		defer fmt.Println(" Exit " , n)
		fmt.Printf("Routine Started  %d\n", n)
		 if rand.Intn(10)<5 {
			return errors.New("Random error")
		 }
		 //...............
		 time.Sleep(10*time.Millisecond)
		 return nil
	}



