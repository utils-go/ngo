package generic

import (
	"container/list"
)

// LinkedListNode represents a node in a LinkedList<T>, equivalent to System.Collections.Generic.LinkedListNode<T>.
type LinkedListNode[T any] struct {
	Value T
	node  *list.Element
}

// LinkedList represents a doubly linked list, equivalent to System.Collections.Generic.LinkedList<T>.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.linkedlist-1?view=netframework-4.7.2
type LinkedList[T any] struct {
	items *list.List
}

// NewLinkedList creates a new empty LinkedList.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		items: list.New(),
	}
}

// NewLinkedListFromSlice creates a new LinkedList from a slice.
func NewLinkedListFromSlice[T any](items []T) *LinkedList[T] {
	ll := NewLinkedList[T]()
	for _, item := range items {
		ll.AddLast(item)
	}
	return ll
}

// Count gets the number of nodes in the LinkedList.
func (ll *LinkedList[T]) Count() int {
	return ll.items.Len()
}

// First gets the first node of the LinkedList.
func (ll *LinkedList[T]) First() *LinkedListNode[T] {
	e := ll.items.Front()
	if e == nil {
		return nil
	}
	return &LinkedListNode[T]{Value: e.Value.(T), node: e}
}

// Last gets the last node of the LinkedList.
func (ll *LinkedList[T]) Last() *LinkedListNode[T] {
	e := ll.items.Back()
	if e == nil {
		return nil
	}
	return &LinkedListNode[T]{Value: e.Value.(T), node: e}
}

// AddFirst adds a new node containing the specified value at the start of the LinkedList.
func (ll *LinkedList[T]) AddFirst(value T) *LinkedListNode[T] {
	e := ll.items.PushFront(value)
	return &LinkedListNode[T]{Value: value, node: e}
}

// AddLast adds a new node containing the specified value at the end of the LinkedList.
func (ll *LinkedList[T]) AddLast(value T) *LinkedListNode[T] {
	e := ll.items.PushBack(value)
	return &LinkedListNode[T]{Value: value, node: e}
}

// AddAfter adds a new node after the specified node.
func (ll *LinkedList[T]) AddAfter(node *LinkedListNode[T], value T) *LinkedListNode[T] {
	e := ll.items.InsertAfter(value, node.node)
	return &LinkedListNode[T]{Value: value, node: e}
}

// AddBefore adds a new node before the specified node.
func (ll *LinkedList[T]) AddBefore(node *LinkedListNode[T], value T) *LinkedListNode[T] {
	e := ll.items.InsertBefore(value, node.node)
	return &LinkedListNode[T]{Value: value, node: e}
}

// Remove removes the specified node from the LinkedList.
func (ll *LinkedList[T]) Remove(node *LinkedListNode[T]) {
	ll.items.Remove(node.node)
}

// RemoveFirst removes the first node.
func (ll *LinkedList[T]) RemoveFirst() {
	f := ll.items.Front()
	if f != nil {
		ll.items.Remove(f)
	}
}

// RemoveLast removes the last node.
func (ll *LinkedList[T]) RemoveLast() {
	b := ll.items.Back()
	if b != nil {
		ll.items.Remove(b)
	}
}

// Clear removes all nodes.
func (ll *LinkedList[T]) Clear() {
	ll.items.Init()
}

// Contains determines whether a value is in the LinkedList.
func (ll *LinkedList[T]) Contains(value T) bool {
	for e := ll.items.Front(); e != nil; e = e.Next() {
		v := e.Value.(T)
		if any(v) == any(value) {
			return true
		}
	}
	return false
}

// Find finds the first node that contains the specified value.
func (ll *LinkedList[T]) Find(value T) *LinkedListNode[T] {
	for e := ll.items.Front(); e != nil; e = e.Next() {
		v := e.Value.(T)
		if any(v) == any(value) {
			return &LinkedListNode[T]{Value: v, node: e}
		}
	}
	return nil
}

// FindLast finds the last node that contains the specified value.
func (ll *LinkedList[T]) FindLast(value T) *LinkedListNode[T] {
	var last *LinkedListNode[T]
	for e := ll.items.Front(); e != nil; e = e.Next() {
		v := e.Value.(T)
		if any(v) == any(value) {
			last = &LinkedListNode[T]{Value: v, node: e}
		}
	}
	return last
}

// ToSlice converts the LinkedList to a slice.
func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.items.Len())
	for e := ll.items.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(T))
	}
	return result
}

// ForEach performs the specified action on each element.
func (ll *LinkedList[T]) ForEach(action func(T)) {
	for e := ll.items.Front(); e != nil; e = e.Next() {
		action(e.Value.(T))
	}
}

// --- LinkedListNode methods ---

// Next gets the next node.
func (n *LinkedListNode[T]) Next() *LinkedListNode[T] {
	if n.node.Next() == nil {
		return nil
	}
	return &LinkedListNode[T]{Value: n.node.Next().Value.(T), node: n.node.Next()}
}

// Previous gets the previous node.
func (n *LinkedListNode[T]) Previous() *LinkedListNode[T] {
	if n.node.Prev() == nil {
		return nil
	}
	return &LinkedListNode[T]{Value: n.node.Prev().Value.(T), node: n.node.Prev()}
}
