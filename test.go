package main

import "fmt"

type num struct{
	x int
	y int
}

func initialize (n *num){
//	n = &num{}
	n.x = 5
	n.y = 4
	fmt.Println("this works")
}

func main (){
	//var x int
	//x = 3
	var y *num;
	//y := &num{}
	
	//n := num {nil, nil}
	if y == nil{
		y = &num {}
		initialize (y)
		fmt.Println("got it")
		//return
	}

	fmt.Println("Hello World!", y.x, y.y)

}	
