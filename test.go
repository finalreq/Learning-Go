package main

import "fmt"

type num struct{
	x int
	y int
}


func main (){
	var x int
	x = 3
	var y *int
	y = &x
	
	n := num {nil, nil}


	fmt.Println("Hello World!", n.x, n.y)

}	
