package main

import "fmt"

type Node struct {
	val int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (T * LinkedList) prepend(node *Node) {
	n := T.head
	T.head = node
	node.next = n
}

func (T * LinkedList) readList() {
	p := T.head
	for p != nil {
		fmt.Println(p.val)
		p = p.next
	}
}

func main() {
	list := LinkedList{}
	node1 := Node{val:4}
	node2 := Node{val:5}
	node3 := Node{val:100}
	node4 := Node{50,nil}
	list.prepend(&node1)
	list.prepend(&node2)
	list.prepend(&node3)
	list.prepend(&node4)
	list.readList()
}