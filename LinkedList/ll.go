package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func add_node(root *Node, nval int) {
	var tmp = root

	n := Node{nval, nil}

	for tmp.next != nil {
		tmp = tmp.next
	}

	tmp.next = &n
	return

}
func print_list(root *Node) {
	var tmp = root

	for tmp != nil {
		fmt.Print(tmp.val, " --> ")
		tmp = tmp.next
	}

	fmt.Println()
	return
}

func remove_node(root **Node, val int) {
	var cur = root.next
	var prev = **root
	for cur.val != val {
		if cur.next == nil {
			fmt.Println("value does not exist in the list")
			return
		}
		cur = cur.next
		prev = *prev.next
	}
	prev.next = cur.next
	prev = cur.next
	cur = cur.next.next
}

func main() {
	var root *Node
	var option int
	var input int

	option = 0

	for true {

		if option == 1 {

			fmt.Print("Please enter a number to enter into the list: ")
			fmt.Scanln(&input)

			if root == nil {
				root = &Node{input, nil}
			} else {
				add_node(root, input)
			}

		} else if option == 2 {
			print_list(root)

		} else if option == 3 {
			fmt.Print("Please enter the number to be removed: ")
			fmt.Scanln(&input)
			remove_node(&root, input)
			//TODO remove number from list

		}

		fmt.Printf("1 ........ add new number \n2 ........ print the list \n3 ........ remove from the list \n")

		fmt.Print("Please enter a new command: ")
		fmt.Scanln(&option)
	}

}
