package main

import "fmt"

func main (){
	var x int
	x = 3
	var y *int
	y = &x
	fmt.Println("Hello World!", x, *y)

}	
