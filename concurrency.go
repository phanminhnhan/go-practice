package main

import (
	"fmt"
	"sync"
	"time"
	//"runtime"
)

var wg = sync.WaitGroup{}


func main() {
	defer timetrack(time.Now())

	wg.Add(2)
	go numbers()
	go alphabets()
	wg.Wait()

	// for i := 1; i <= 20000; i++ {
	// 		 fmt.Println(i)
	// 	 }
		 
	
	// for i := 1; i <= 20000; i++ {
	// 	fmt.Println("aaaaa")
	// }	 
		
}


func timetrack(starttime time.Time){
	showtime := time.Since(starttime)
	fmt.Println("took", showtime)
}

func numbers() {
	for i := 1; i <= 20000; i++ {
		go fmt.Println(i)
	}

	wg.Done()
}


func alphabets() {
	for i := 1; i <= 20000; i++ {
		fmt.Println("aaaa")

	}

	wg.Done()
}