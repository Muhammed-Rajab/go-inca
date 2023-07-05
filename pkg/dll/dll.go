// Change this to dll
package main

import "fmt"

// ////////////////////////////
// NODE
// ////////////////////////////
type Node struct {
	val  string
	prev *Node
	next *Node
}

// Function to create a node
func CreateNode(val string, prev, next *Node) *Node {
	return &Node{val, prev, next}
}

// Method to display a node
func (node *Node) Display() {
	fmt.Printf("value: %s\n", node.val)
}

// Method to delete a node
func (node *Node) Delete() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

// ////////////////////////////
// DOUBLY LINKED LIST
// ////////////////////////////
type DoublyLinkedList struct {
	head *Node
}

// Function to create a DLL
func CreateDoublyLinkedListEmpty() *DoublyLinkedList {
	return &DoublyLinkedList{nil}
}

func CreateDoublyLinkedListWithHead(headVal string) *DoublyLinkedList {
	head := CreateNode(headVal, nil, nil)
	return &DoublyLinkedList{head}
}

// Method to display the dll
func (dll *DoublyLinkedList) Display() {
	if temp := dll.Head(); temp != nil {
		fmt.Printf("nil ⇄ ")
		for temp != nil {
			fmt.Printf("%s ⇄ ", temp.val)
			temp = temp.next
		}
		fmt.Printf("nil\n")
	}

}

// Method to get length of the dll
// Better way: keep track of the length without calculating it every single time
func (dll *DoublyLinkedList) Length() uint32 {
	len := 0
	temp := dll.Head()
	for temp != nil {
		len++
		temp = temp.next
	}
	return uint32(len)
}

// Method to get Head Node
func (dll *DoublyLinkedList) Head() *Node {
	return dll.head
}

// Method to get Tail Node
func (dll *DoublyLinkedList) Tail() *Node {
	temp := dll.Head()
	for temp.next != nil {
		temp = temp.next
	}
	return temp
}

// Method to set Head Node
func (dll *DoublyLinkedList) setHead(node *Node) *Node {
	dll.head = node
	node.next = nil
	node.prev = nil
	return node
}

// Method to append a node to the end of the dll
func (dll *DoublyLinkedList) Append(node *Node) *Node {
	// Case: 1 -> Head is nil
	if dll.Head() == nil {
		return dll.setHead(node)
	}

	// Case: 2 -> Head is not nil
	tail := dll.Tail()
	tail.next = node
	node.prev = tail
	node.next = nil
	return node
}

// Method to prepend a node to the beginning of the dll
func (dll *DoublyLinkedList) Prepend(node *Node) *Node {
	// Case: 1 -> Head is nil
	if dll.Head() == nil {
		return dll.setHead(node)
	}

	// Case: 2 -> Head is not nil
	node.next = dll.head
	dll.head.prev = node
	dll.head = node
	node.prev = nil
	return node
}

// Method to remove from the last element
func (dll *DoublyLinkedList) Pop() *Node {
	if dll.Head() == nil {
		return nil
	}

	tail := dll.Tail()
	if tail.prev != nil {
		tail.prev.next = nil
	} else {
		dll.head = nil
	}
	tail.next = nil
	return tail
}

func main() {
	dll := CreateDoublyLinkedListEmpty()
	dll.Prepend(CreateNode("nigga", nil, nil))
	dll.Prepend(CreateNode("bro", nil, nil))
	dll.Append(CreateNode("2", nil, nil))
	dll.Append(CreateNode("4", nil, nil))
	dll.Append(CreateNode("3", nil, nil))
	dll.Pop()
	dll.Pop()
	dll.Pop()
	dll.Pop()
	dll.Pop()
	dll.Pop()
	dll.Append(CreateNode("sdf", nil, nil))
	dll.Display()
}
