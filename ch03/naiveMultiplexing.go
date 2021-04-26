package main

import "fmt"


func main(){
	channels := [5](chan int){
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go func(){
		for _, chx := range channels {
			fmt.Println("Receiving from", <-chx)
		}
	}()

	for i := 1; i < 6; i++{
		fmt.Println("Sending on channel:", i)
		channels[i] <- i
	}
}
