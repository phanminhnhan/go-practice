package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*var object = map[string ]interface{}{
		"id ": 1,
		"name":"phan",
		"class":"43k21",
	}

	// chuyen doi du lieu tu Map sang json

	data1, err := json.Marshal(object)

	if err == nil {
		//convert to string
		var mydata = string(data1)

		fmt.Println(mydata)
	}

	// chuyen doi tu json sang map

	var playload = `{
		"id": 1 ,
		"name": "phan minh nhan",
		"class": "43k21"
	}`
	var data2 = map[string]interface{}{}
	err1 := json.Unmarshal([]byte(playload), &data2)
	if err1 == nil {
		fmt.Println(data2)
	}else{
		panic(err1)
	}
*/


	var playload2 string = `{
		"age": 1 ,
		"name": "phanminhnhan"
	}`

	doituong:= Object{}

	err := json.Unmarshal([]byte(playload2), &doituong)
	if err != nil {
		panic(err)
	}

	fmt.Println(doituong)

}

type Object struct{
	Age  int 	`json:"age"`
	Name string `json:"name"`
}
