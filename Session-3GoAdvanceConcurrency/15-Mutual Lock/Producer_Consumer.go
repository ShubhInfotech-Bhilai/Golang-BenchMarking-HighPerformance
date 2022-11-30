package main
import "fmt"
var done = make (chan bool)
	var msgs = make(chan int)

func produce(){
	for i:=0;i<=10;i++{
		    fmt.Println(" Send")
			msgs<-i
		}
		done<-true

	}

func consume(){
	
	for{
		
		fmt.Println("Recieved ")
		msg:= <-msgs
		fmt.Println(msg)
		
	}
	
}	
func main(){
	go produce()
	go consume()
	<-done
	
	
}