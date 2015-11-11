// This will be a simple bst implementation in go

package main

import "fmt"

type Node struct {
	val   int
	right *Node
	left  *Node
}

func add_node(root *Node, input int) {

	if input > root.val {
		if root.right == nil {
			root.right = &Node{input, nil, nil}
			return
		}
		add_node(root.right, input)

	} else if input < root.val {
		if root.left == nil {
			root.left = &Node{input, nil, nil}
			return
		}
		add_node(root.left, input)
	} else {
		fmt.Println("Value is already in the tree!")
		return
	}

}

func print_inorder(root *Node) {
	tmp := root
	if tmp == nil {
		return
	}
	print_inorder(root.left)
	fmt.Printf("%d ", root.val)
	print_inorder(root.right)
}

func remove_node(root **Node, input int) {

}

func find_node(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if root.val == val {
		return true
	}
	if root.val > val {
		return find_node(root.left, val)
	} else {
		return find_node(root.right, val)
	}
}

func main() {
	var root *Node
	var option int
	var input int
	var found bool
	for true {
		fmt.Printf("1 ........ add new number\n2 ........ print the tree in order\n3 ........ remove from the tree\n4 ........ Find the number in the tree\n")
		fmt.Scanln(&option)
		if option == 1 {
			fmt.Print("Please enter a number to be entered into the list: ")
			fmt.Scanln(&input)
			if root == nil {
				root = &Node{input, nil, nil}
			} else {
				add_node(root, input)
			}
		} else if option == 2 {
			//for now only do in order, will add more later
			print_inorder(root)
			fmt.Printf("\n")
		} else if option == 3 {
			fmt.Print("Please enter the number you want to remove from the tree")
			fmt.Scanln(&input)
			remove_node(&root, input)
		} else if option == 4 {
			fmt.Print("Pleaase enter the number you want to search the tree for: ")
			fmt.Scanln(&input)
			found = find_node(root, input)
			if found == true {
				fmt.Println("The value exists in the tree")
			} else {
				fmt.Println("The value does not exist in the tree")
			}
		}
	}

	fmt.Println("Hello World!")

}
