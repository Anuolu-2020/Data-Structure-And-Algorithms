package main

import "fmt"

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}

type DoubleLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (dl *DoubleLinkedList) Insert(value interface{}) {
	newNode := &Node{Data: value}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		newNode.Previous = dl.Tail
		dl.Tail.Next = newNode
		dl.Tail = newNode
	}

	dl.Length++
}

func (dl *DoubleLinkedList) InsertAtBegin(value interface{}) {
	newNode := &Node{Data: value}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		dl.Length++
	} else {
		newNode.Next = dl.Head
		dl.Head.Previous = newNode
		dl.Head = newNode
		dl.Length++
	}
}

func (dl *DoubleLinkedList) InsertAtIndex(value interface{}, index int) {
	newNode := &Node{Data: value}

	if index == 0 {
		if dl.Head == nil {
			dl.Head = newNode
		} else {
			newNode.Next = dl.Head
			dl.Head.Previous = newNode
			dl.Head = newNode
		}
		dl.Length++
		return
	}

	current := dl.Head
	i := 0
	for i < index && current != nil {
		current = current.Next
		i++
	}

	if current == nil {
		tail := dl.Head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = newNode
		newNode.Previous = tail
	} else {
		newNode.Previous = current.Previous
		newNode.Next = current
		if current.Previous != nil {
			current.Previous.Next = newNode
		}
		current.Previous = newNode
	}

	dl.Length++
}

func (dl *DoubleLinkedList) Delete() {
	current := dl.Head

	if current == nil {
		return
	}

	if dl.Length == 1 {
		dl.Head = nil
		dl.Tail = nil
		dl.Length--
	}

	indexBeforeEnd := dl.Length - 1

	for i := 1; i < indexBeforeEnd; i++ {
		current = current.Next
	}

	current.Next = nil
	dl.Tail = current

	dl.Length--
}

func (dl *DoubleLinkedList) Print() {
	current := dl.Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (dl *DoubleLinkedList) PrintReverse() {
	current := dl.Tail

	for current != nil {
		fmt.Println(current.Data)
		current = current.Previous
	}
}

func (dl *DoubleLinkedList) PrintLen() {
	fmt.Printf("Double linked list length is: %d\n", dl.Length)
}

func main() {
	doubleLinkedList := &DoubleLinkedList{}

	doubleLinkedList.Insert(20)
	doubleLinkedList.Insert(30)
	doubleLinkedList.Insert("Manuel")
	doubleLinkedList.Insert(90)
	doubleLinkedList.Insert("Madrid")
	doubleLinkedList.InsertAtBegin(2000)
	doubleLinkedList.Delete()
	doubleLinkedList.InsertAtIndex("Index", 2)

	doubleLinkedList.Print()
	fmt.Println()
	fmt.Println("Printing Backwards")
	doubleLinkedList.PrintReverse()

	doubleLinkedList.PrintLen()
}

