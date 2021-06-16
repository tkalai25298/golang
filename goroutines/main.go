package main

import (
	"fmt"
	"time"
)

func printString(s string) {
	fmt.Println(s)
}

func main() {
	fmt.Println("main go-routine started")

	// printString("Go-routine")
	go printString("Go-routine") //wont print coz main routine wont block..moves ctrl to next line

	time.Sleep(2*time.Millisecond)
	
	fmt.Println("main go-routine ended")
}
