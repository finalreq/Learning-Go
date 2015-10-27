package main

import "fmt"

type Node struct{
	val int
	next *Node
}

func add_node (root *Node){
	var tmp = root
	tmp.val = 2

}

func main (){
	root := Node{}
	root.val = 1

	fmt.Println("Hello world!", root.val)

	add_node (&root)

	fmt.Println(root.val)
	
}
