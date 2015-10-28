package main

import "fmt"

type Node struct{
	val int
	next *Node
}

func add_node (root *Node, nval int){
	if root.val == nil{
		root.val = nval
		return
	}
	var tmp = root
	n := Node {nval, nil}
	for tmp.next != nil{
		tmp = tmp.next
	}
	tmp.next = &n
	return

}
func print_list(root *Node){
	var tmp = root
	for tmp.next != nil{
		fmt.Print(tmp.val, " --> ")
		tmp = tmp.next
	}
	fmt.Println()
	return
}


func main (){
	root := Node{}
	var option int
	var input int
//	fmt.Print("Please enter a number to enter into the list: ")
//	fmt.Scanln(&input);
//	root.val = input
//	root.next = nil
	option = 0
	for true{
		if option == 1{
			fmt.Print("Please enter a number to enter into the list: ")
			fmt.Scanln(&input);
			add_node(&root, input)
		
		} else if option == 2{
			print_list(&root)
		
		} else if option == 3{
			//remove number from list
		
		} 

		fmt.Printf("1 ........ add new number \n2 ........ print the list \n3 ........ remove from the list \n")

		fmt.Print("Please enter a new command: ")
		fmt.Scanln(&option)
	}
	
}
