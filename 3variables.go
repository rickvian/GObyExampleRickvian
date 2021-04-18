package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	fmt.Println("type of a:", reflect.ValueOf(a).Kind())

	fmt.Println("type of a: %T\n", a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println("type of: ", reflect.ValueOf(b).Kind(), ", type of: ", reflect.ValueOf(c).Kind())

	var d = true
	fmt.Println(d)
	fmt.Println(reflect.ValueOf(d).Kind())

	var e int
	fmt.Println(e)
	fmt.Println("type of: ", reflect.ValueOf(e).Kind())

	f := "apple"
	fmt.Printf("f is of type %T\n", f)
	fmt.Println("type of: ", reflect.ValueOf(f).Kind())
}
