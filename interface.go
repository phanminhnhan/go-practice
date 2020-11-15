

package main

import "fmt"


func main() {

	uni := &Student( "uni", "43k21")
	fmt.Println(implementinterface(uni))

}




func implementinterface(d descriptor)string {
	return fmt.Sprintf("Dear %s", d.description())
}

type descriptor interface{
	description()
}


type Student struct{

	name string
	class string
}

func(s *Student) description()string{

	return fmt.Sprintf("%s %s", s.name, s.class)
}