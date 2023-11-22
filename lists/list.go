package lists

import "github.com/246859/containers"

// List is the base interface of all lists
type List[T any] interface {
	// Get returns the elem at the given index
	Get(i int) (T, bool)
	// IndexOf returns the index of the given element, if not found, return -1
	IndexOf(elem T) int
	// Set replaces the elem at the given index
	Set(i int, elem T)
	// Add adds given elems to the list
	Add(elem ...T)
	// Insert inserts the given elems into the list at the given index
	Insert(i int, elem ...T)
	// Remove removes the elem match given index from the list
	Remove(i int)
	// RemoveElem removes the given element from the list
	RemoveElem(elem T)
	// Contains check if list contains the given elements
	Contains(elem T) bool
	// Clone returns a copy of the original list
	Clone() List[T]
	// Join joins the given list into the original list
	Join(list List[T])
	// Iterator returns a list iterator
	Iterator(reverse bool) ListIterator[T]

	containers.Container[T]
}

type ListIterator[T any] interface {
	containers.Iterator[int, T]
}