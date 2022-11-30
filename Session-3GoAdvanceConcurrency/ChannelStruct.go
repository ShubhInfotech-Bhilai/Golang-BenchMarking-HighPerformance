package main

import "fmt"
type Group struct{
	ServerPort int 
	IP string 
	n int
}
func main(){
	c:=make(chan Group)
	go func (){
		t:=Group{}
		for i:=0;i<=10;i++{
            t.ServerPort=8080
			t.IP="192.0.0.1"
			t.n=i
			c<- t

		}
		
		close(c)
	}()

	for j:= range c{
		fmt.Println( j)
	} 


}