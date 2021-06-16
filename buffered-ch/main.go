package main

import "fmt"


func fibonacci(ch chan int) {
	x,y := 0,1
	n := cap(ch)
	for i:= 0; i<n; i++ {
		ch <- x
		x,y = y,x+y
	}
	close(ch)
}
func main() {
	ch := make(chan int,10)
	go fibonacci(ch)
	
	for i := range ch {
		fmt.Println(i)
	}
}