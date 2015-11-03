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
	for cur != nil {
		if (*prev).val == val {
			*prev = (*prev).next
			cur = cur.next
			tmp = (*prev)
			continue
		}
		if (*cur).val == val {
			(*prev).next = cur.next
			*prev = cur.next

			if cur.next == nil {
				break
			}

			cur = cur.next.next
			continue
		}

		cur = cur.next
		*prev = (*prev).next
	}

	*root = tmp
}

func find_node(root *Node, val int) {
	count := 0
	for root != nil {
		if root.val == val {
			count++
		}
		root = root.next
	}
	if count > 0 {
		fmt.Println("\nThere are ", count, " occurences of value ", val, " in the list\n")
	} else {
		fmt.Println("\nValue ", val, " does not exist in the list\n")
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

		} else if option == 4 {
			fmt.Print("Please enter the number you are looking for ")
			fmt.Scanln(&input)
			find_node(root, input)
		}

		fmt.Printf("1 ........ add new number \n2 ........ print the list \n3 ........ remove from the list \n4 ........ Find the number of occurences of a value in the list\n")

		fmt.Print("Please enter a new command: ")
		fmt.Scanln(&option)
	}

}
