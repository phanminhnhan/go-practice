package main


import(
	"fmt"
)

func main(){

	
	// for i := 0; i <= 10  ; i++{
	// 	fmt.Println(fib(i))
	// }


	fmt.Println(fib(3))

}

func fib (i int) int{
	if i <= 1 {
		return i
	}
	return fib(i-1)+ fib(i-2)
}