package main

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] interface {
	Insert(idx int, value T)
	Remove(value T)
	Push(T)                        //completed
	Pop() (value T, err error)     // completed
	dequeue() (value T, err error) // completed
	Reverse()
	Search(value T) (T, error)
	Print() //completed
}

func newLinkedList[T comparable](value T) *LinkedListImpl[T] {
	return &LinkedListImpl[T]{head: &Node[T]{Value: value}}
}

type LinkedListImpl[T comparable] struct {
	head *Node[T]
}

func (N *LinkedListImpl[T]) Print() {
	current := N.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (N *LinkedListImpl[T]) Push(value T) {
	current := N.head
	if current == nil {
		panic(fmt.Sprintf("cannot push to an uninitialized list: attempted to push value %v", value))
	}
	for current != nil {
		if current.Next == nil {
			current.Next = &Node[T]{Value: value}
			break
		}
		current = current.Next
	}
}

func (N *LinkedListImpl[T]) dequeue() (T, error) {
	current := N.head
	if current == nil {
		var zero T
		return zero, errors.New("the list is already empty")
	}

	value := current.Value
	N.head = N.head.Next
	return value, nil

}

func (N *LinkedListImpl[T]) Pop() (T, error) {
	current := N.head
	if current == nil {
		var zero T
		return zero, errors.New("the list is already empty")
	}
	var value T
	var prev *Node[T]
	for current != nil {
		if current.Next == nil {
			value = current.Value
			if prev != nil {
				prev.Next = nil
			}
			break
		}
		prev = current
		current = current.Next
	}
	return value, nil
}

func main() {
	linkedList := newLinkedList(10)

	linkedList.Push(6)
	linkedList.Push(10)

	fmt.Println("printed by print")
	linkedList.Print()

	linkedList.Pop()

	fmt.Println("printed by print")
	linkedList.Print()

}
