package main

import "testing"

func TestDoubleLinkedList(t *testing.T) {
	doubleLinkedList := &DoubleLinkedList{}

	doubleLinkedList.Insert(20)
	doubleLinkedList.Insert(30)
	doubleLinkedList.Insert("Manuel")
	doubleLinkedList.Insert(90)
	doubleLinkedList.Insert("Madrid")
	doubleLinkedList.InsertAtBegin(2000)
	doubleLinkedList.Delete()
	doubleLinkedList.InsertAtIndex("Index", 2)

	correctLinkedListData := []interface{}{2000, 20, "Index", 30, "Manuel", 90}

	current := doubleLinkedList.Head

	i := 0
	for current != nil {
		if current.Data != correctLinkedListData[i] {
			t.Errorf(
				"tests[%d] - linked-list data wrong, expected=%v, got=%v",
				i,
				current.Data,
				correctLinkedListData[i],
			)
		}
		current = current.Next
		i++
	}
}
