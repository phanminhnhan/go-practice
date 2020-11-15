package main

import "fmt"

func main() {
	var a = map[interface{}]interface{}{
		"nhan":"male", 1 :"female",
	}
	b := a
	fmt.Println(a,b)
	// change value
	fmt.Println("=========changed==========")
	b["trang"]= "FEMALE"
	// add one more pair
	a["vy"] = 19
	fmt.Println("=========added=============")
	fmt.Println(a,b)
	// update value
	fmt.Println("=========updated=============")
	a[1] = "name"
	fmt.Println(a,b)
	// check exist
	_, isExist := a["bao"]
	fmt.Println(isExist)

}
