package main 

import "fmt"

type Node struct {
	val int
	next * Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addNodes(p * Node) {
	p.next = list.head
	list.head  = p
}

func (list * LinkedList) printNodes() {
	itr := list.head
	for itr != nil {
		fmt.Println(itr.val)
		itr = itr.next
	}
}

func main() {
	node1 := Node{val:1}
	node2 := Node{val:2}
	node3 := Node{val:3}
	node4 := Node{val:4}

	list := LinkedList{}
	list.addNodes(&node1)
	list.addNodes(&node2)
	list.addNodes(&node3)
	list.addNodes(&node4)

	list.printNodes()
}