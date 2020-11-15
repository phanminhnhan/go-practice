package main

import (
	"fmt"
)

type animal struct {
	id int
	name string
}


func (a animal) F (ten string){
	a.name = ten
}

func (a *animal) PF (ten string) {
	a.name = ten
}



func main() {
	// a := &animal{1,"fesafvsadf"}
	// fmt.Println(a)
	// a.PF("nhan")
	// fmt.Println(a)


	a := 1
	b := &a 
	*b = 1000

	fmt.Println(a)
}
