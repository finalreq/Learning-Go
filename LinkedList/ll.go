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
	//This first if case works
	if (*root).next == nil { //If there is only one element in the list
		if (*root).val == val {
			(*root) = nil
			return
		} else {
			fmt.Println("Number is not in the list")
			return
		}
	}

	cur := (*root).next
	tmp := (*root)
	prev := root
	_ = tmp
	for cur.val != val {
		fmt.Println("loop")
		if (*cur).next == nil && (*cur).val != val { // if at the end of the list
			fmt.Println("value does not exist in the list")
			return
		} else if (*cur).next == nil && (*cur).val == val {
			(*prev).next = nil
			return
		}
		cur = cur.next
		//fmt.Println("We get here")
		*prev = (*prev).next
		//fmt.Println("we go there", (*prev).val)
	}

	(*prev).next = cur.next
	*prev = cur.next
	(*root) = tmp
	if cur.next == nil {
		return
	} else {
		cur = cur.next.next
	}
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
