package main
import (
	//"fmt"
	"sync"
	"testing"
)

func addByShareMemory(n int )[]int{

	var ints [] int 
	var wg sync.WaitGroup 
	var mux sync.Mutex 
	wg.Add(n)  // wating pools 
	 for i:=0;i<n ;i++{
		  	go func (i int){
				defer wg.Done()
				mux.Lock()
				ints=append(ints, i)
				mux.Unlock()		
               
			}(i)		
	 }
	 wg.Wait()

	 return ints
}

func addByShareCommunication(n int )[]int{
	var ints [] int
	channel :=make(chan int,n) 
	for i:=0;i<n ;i++{
		go func (channel chan<- int , value int){
			channel<-value
		} (channel,    i)
		}
     for i:= range channel{
		ints=append(ints, i)
		if len(	ints) ==n {
			break
		}
	 }

	 close(channel)
	 return ints

}
func BenchmarkAddByShareCommunication(b *testing.B){

	for i:=0;i<b.N ;i++{ 
		addByShareCommunication(25000)
	}
}

func BenchmarkAddByShareMemory(b *testing.B){

	for i:=0;i<b.N ;i++{
		addByShareMemory(25000)
	 }
}

/*

func main(){
	asbm:=addByShareMemory(10)
	fmt.Println(len(asbm)   ,  asbm)
    

	asbc:=addByShareCommunication(10)
	fmt.Println(len(asbc)   ,  asbc)

}*/