package main


import (
	"fmt"
	"sync"
)


func main(){
	var wg sync.WaitGroup

	ordersProcessed := 0
	cashier := func(orderNum int){
		defer wg.Done()
		if ordersProcessed < 10 {
			fmt.Println("Processing order", orderNum)
			ordersProcessed++
		} else {
			fmt.Println("I am tried! I want to take rest!", orderNum)
		}
	}

	for i := 0; i < 30; i ++ {
		wg.Add(1)
		go func(orderNum int){
			cashier(orderNum)
		}(i)
	}
	wg.Wait()
}
