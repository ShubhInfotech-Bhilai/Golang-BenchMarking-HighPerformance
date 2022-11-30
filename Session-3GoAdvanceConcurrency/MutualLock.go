package main
import (
	"fmt"
)
func main() {
	exit:=make(chan struct{})
	transfer:=make(chan int)
	go func(){
		defer func(){ 	close(exit) }()

		for i:=0;i<=20; i++{
			transfer<-i
			fmt.Println("Tast Trasfer from T1() -T2()\t")
		}
	}()
	go func(){
		for i:=0;i<=20;i++{
            fmt.Println(" Recieved T2() \t" ,<-transfer)
		}
	}()
<-exit
}
