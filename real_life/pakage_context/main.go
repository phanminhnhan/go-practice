package main

import (
	"context"
	"time"
	"fmt"
	
)


func main(){
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)

	doSomething(ctx)

	// time.AfterFunc(time.Microsecond , func(){
	// 	cancel()
	// })
	// defer cancel() // giup dam bao qua trinh da duwoc cancel
	
}

func doSomething(ctx context.Context){
	
	canceled := make(chan bool)


	go func() {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			canceled <- true 
			return
			
		}
	}()

	isCanceled := <- canceled
	if (isCanceled){
		close(canceled)
		return
	}
	go func(){
		fmt.Println("Start...")
		fmt.Println("end...")
	}()
}