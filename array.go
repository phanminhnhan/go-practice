package main

import "fmt"

func main() {
	// ar := []int{1,2,3,4}

	// ar = append(ar, 4,6,7,878)
	// for _, i := range(ar){
	// 	if i %2 == 0{
	// 		fmt.Println(i, "la so chan ")
	// 	}else {
	// 		fmt.Println(i, "la so le")
	// 	}
	// }

	//ar := [3]int{1,2,3}
	//fmt.Println(ar)
	//b := &ar
	//fmt.Println(b)
	//b[2] = 100
	//fmt.Println(b)
	//fmt.Println(ar)
	//change(ar[:])
	//fmt.Println(ar)
	a := []int{1,2,3,4,5,6,7,8,9}

	b := a
	b[2] = 100
	fmt.Println(a)
	for i, v  := range a {
		fmt.Println("index:", i , "v:" ,v)
	}
}

//func change(sls []int) {
//	sls[0] = 100
//}