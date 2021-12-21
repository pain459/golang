package main // use of package keyword

import (
	"fmt"
	"time"
) // use of import keyword

type person struct {  // use of type keyword
	name string
	age int
}  // use of struct keyword

type interface_example interface{
	interface_function_int() int
	interface_function_str() string	
}  // use of interface keyword

func interface_function_str() (string){
	return "returning from function"  // use of return keyword
}

func main() {  // func keyword
	var str = "hello world"  // var keyword
	const num = 10 // const keyword
	arr := [6] int {10, 20, 30, 40, 50, 60}
	map_data := map[string]int{
		"hello":80,
		"world":45,
	}  // declare map
	cs1 := make(chan string)  // declare channel
	cs2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		cs1<- "one"
	} ()
	fmt.Println(str, num, map_data, cs1)
	if num > 100{ // use of if keyword
		fmt.Println("num is greater than 100")
	} else {  // use of else keyword
		fmt.Println("num is not greater than 100")
	}
	for pos, val := range arr {  // use of for and range keyword
		if (pos == 5){
			break  // break keyword
		}
		if (pos == 1){
			continue  // continue keyword
		}
		fmt.Printf("value at index %d is %d\n", pos, val)
	}
	fmt.Println(interface_function_str())
	switch num {  // switch keyword
	case 1: // use of case keyword
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: // default keyword
		fmt.Println("No match")
	}
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 10:
		fmt.Println("three")
	fallthrough // use of fallthrough keyword
	default:
		fmt.Println("default statement")
	}
	defer interface_function_str()  // use of defer keyword
	go interface_function_str()  // use of go keyword
	select {
	case msg1 := <-cs1:
		fmt.Println("received", msg1)
	case msg2 := <-cs2:
		fmt.Println("received", msg2)
	}  // use of select keyword
	goto gotolabel  // use og goto keyword
	gotolabel:
	fmt.Println("Inside got statement")
}