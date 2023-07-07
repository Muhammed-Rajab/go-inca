// Change this to dll
package dll

import (
	"fmt"
	"time"
)

// ////////////////////////////
// NODE
// ////////////////////////////
type Node struct {
	// Data
	Key      string
	Val      string
	TTL      time.Duration
	StoredAt time.Time

	// Pointers
	Prev *Node
	Next *Node
}

// Function to create a node
func CreateNode(key, val string, ttl time.Duration, prev, next *Node) *Node {
	return &Node{
		key, val,
		0 * time.Millisecond,
		time.Now(),
		prev, next,
	}
}

// Method to display a node
func (node *Node) Display() {
	fmt.Printf("value: %s\n", node.Val)
}

// Method to delete a node
func (node *Node) Delete() {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

// ////////////////////////////
// DOUBLY LINKED LIST
// ////////////////////////////
type DoublyLinkedList struct {
	HeadNode *Node
	TailNode *Node
}

// Function to create a DLL
func CreateDoublyLinkedListEmpty() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil}
}

func CreateDoublyLinkedListWithHead(headKey, headVal, tailKey, tailVal string) *DoublyLinkedList {
	head := CreateNode(headKey, headVal, -1, nil, nil)
	tail := CreateNode(tailKey, tailVal, -1, nil, nil)
	return &DoublyLinkedList{head, tail}
}

// Method to display the dll
func (dll *DoublyLinkedList) Display() {
	if temp := dll.Head(); temp != nil {
		fmt.Printf("nil ⇄ ")
		for temp != nil {
			fmt.Printf("%s ⇄ ", temp.Val)
			temp = temp.Next
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
		temp = temp.Next
	}
	return uint32(len)
}

// Method to get Head Node
func (dll *DoublyLinkedList) Head() *Node {
	return dll.HeadNode
}

// Method to get Tail Node
func (dll *DoublyLinkedList) Tail() *Node {
	// Implementation 1 -> O(n)
	// temp := dll.Head()
	// for temp.Next != nil {
	// 	temp = temp.Next
	// }
	// return temp
	// Implementation 2 -> O(1), supposedly :)
	return dll.TailNode
}

// Method to set Head Node
func (dll *DoublyLinkedList) setHead(node *Node) *Node {
	dll.HeadNode = node
	dll.TailNode = node
	node.Next = nil
	node.Prev = nil
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
	tail.Next = node
	node.Prev = tail
	node.Next = nil
	dll.TailNode = node
	return node
}

// Method to prepend a node to the beginning of the dll
func (dll *DoublyLinkedList) Prepend(node *Node) *Node {
	// Case: 1 -> Head is nil
	if dll.Head() == nil {
		return dll.setHead(node)
	}

	// Case: 2 -> Head is not nil
	node.Next = dll.HeadNode
	dll.HeadNode.Prev = node
	dll.HeadNode = node
	node.Prev = nil
	return node
}

// Method to remove from the last element
func (dll *DoublyLinkedList) Pop() *Node {
	if dll.Head() == nil {
		return nil
	}

	tail := dll.Tail()
	if tail.Prev != nil {
		tail.Prev.Next = nil
		dll.TailNode = tail.Prev
	} else {
		dll.HeadNode = nil
		dll.TailNode = nil
	}
	tail.Next = nil
	return tail
}
