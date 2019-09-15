package main

import "fmt"

type Node struct {
	value int
	head bool
	next *Node
	size int
}

// SLList stands for Singly-Linked List
func SLList(value int) *Node {
	return &Node{
		value,
		true,
		nil,
		1,
	}
}

func (n *Node) appendNode(value int) *Node {

	n.size += 1
	node := &Node{
		value,
		false,
		nil,
		n.size,
	}

	if n.next == nil {
		n.next = node
	} else {
		curr := n
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node

	}

	return n
}

func (n *Node) printList() {
	curr := n
	for curr.next != nil {
		fmt.Println(curr)
		curr = curr.next
	}
		fmt.Println(curr)
}

func (n *Node) unshift() *Node {
	n.size--
	node := &Node{
		n.next.value,
		true,
		n.next.next,
		n.size,
	}

	return node
}

func main() {
	list := SLList(3)
	_ = list.appendNode(55)
	_ = list.appendNode(44)

	fmt.Printf("%T\n", list)
	list.printList()
	newlist := list.unshift()
	newlist.printList()

}
