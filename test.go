package main

import "fmt"

type num struct {
	x    int
	next *num
}

func initialize(n **num) {
	//	n = num{}
	//n.x = 5
	//n.y = 4
	var tmp *num
	//tmp := &num{}
	tmp = *n
	//_ = tmp
	//**n = *n.next
	tmp = tmp.next
	fmt.Println("this works", tmp.x)
}

func main() {
	//var x int
	//x = 3

	var y *num
	var x *num
	y = &num{5, nil}
	x = &num{4, nil}

	//n := num {nil, nil}
	y.next = x

	fmt.Println(y.x)

	initialize(&y)

	fmt.Println(y.x)

}
