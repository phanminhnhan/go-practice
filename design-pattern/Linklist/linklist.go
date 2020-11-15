package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func makeLL(n *Node) {
	for i := 1; i <= 30; i++ {
		n.value = i
		n.next = new(Node)
		n = n.next
	}
}

func show(n *Node) {
	for n.next != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func main() {
	head := &Node{}
	makeLL(head)
	show(head)
}
