package main


import (
	"fmt"
	"sync"
)

func createCashier(cashierID int, wg *sync.WaitGroup)func(int){
	ordersProcessed := 0
	return func(orderNum int){
		defer wg.Done()
		if ordersProcessed < 10 {
			fmt.Println(cashierID, "->", ordersProcessed)
			ordersProcessed++
		}else{
			fmt.Println("Cashier ", cashierID, "I am tired! I want to take rest!", orderNum)
		}
	}
}


func main(){
	cashierIndex := 0
	var wg sync.WaitGroup

	cashiers := []func(int){}
	for i := 1; i <= 3; i ++ {
		cashiers = append(cashiers, createCashier(i, &wg))
	}

	for i := 0; i < 30; i ++ {
		wg.Add(1)

		cashierIndex = cashierIndex % 3

		func(cashier func(int), i int){
			go cashier(i)
		}(cashiers[cashierIndex], i)

		cashierIndex ++
	}

	wg.Wait()
}
