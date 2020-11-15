package main

import "fmt"

func main() {

	fmt.Println("===simple function===")
	q, _:= showmess(3, 6)
	fmt.Println(q)


	fmt.Println("===multiple parameter function===")

	multiple(234,53453,"ửqwrwreqwrqwr")

	fmt.Println("===Anonymous function===")
	showdata(caculate(2,3))


}

//simple function

func showmess(msg, a int ) ( b int, err error) {

	b = msg * a
	return b, err
}



// multiple parameter function

func multiple  (a ...interface{}) {
	fmt.Println(a)
}

// Anonymous functions



var (
	showdata = func(data interface{}){
		fmt.Println(data)
	}


	caculate = func(a,b int) int {
		return a + b
	}
)

