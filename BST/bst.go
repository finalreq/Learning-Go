// This will be a simple bst implementation in go

package main

import "fmt"

type Node struct {
	val   int
	right *Node
	left  *Node
}

func add_node(root Node, input int) {

}

func print_inorder(root Node) {

}

func remove_node(root *Node, input int) {

}

func find_node(root Node) {

}

func main() {
	var root *Node
	var option int
	var input int

	for true {
		fmt.Printf("1 ........ add new number\n2 ........ print the tree in order\n3 ........ remove from the tree\n4 ........ Find the number in the tree\n")
		fmt.Scanln(&option)
		if option == 1 {
			fmt.Print("Please enter a number to be entered into the list: ")
			fmt.Scanln(&input)
			if root == nil {
				root = &Node{input, nil, nil}
			} else {
				add_node(*root, input)
			}
		} else if option == 2 {
			//for now only do in order, will add more later
			print_inorder(*root)
		} else if option == 3 {
			fmt.Print("Please enter the number you want to remove from the tree")
			fmt.Scanln(&input)
			remove_node(root, input)
		} else if option == 4 {
			fmt.Print("Pleaase enter the number you want to search the tree for: ")
			fmt.Scanln(&input)
			find_node(*root)
		}
	}

	fmt.Println("Hello World!")

}
