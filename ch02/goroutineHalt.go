package main


import (
	"fmt"
	"sync"
)


func simpleFunc(index int, wg *sync.WaitGroup){
        defer wg.Done()
        fmt.Println("Attempting x/(x-10) where x = ", index, " answer is : ", index/(index-10))
}


func main(){
	var wg sync.WaitGroup
        wg.Add(40)
        for i := 0; i < 40; i += 1 {
                go func(j int) {
                        simpleFunc(j, &wg)
                }(i)
        }

        wg.Wait()

}
