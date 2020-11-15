package main

import (
	"fmt"
	"sync"
)


func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	wg.Add(2)

	go func(ch <- chan int ){
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	// chuyen gui vao channel
	go func(ch chan <- int){
		ch <- 5
		ch <- 10
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

}
