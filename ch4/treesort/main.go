package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	myTree := add(nil, 11)
	add(myTree, 5)
	add(myTree, 3)
	add(myTree, 9)
	add(myTree, 20)
	add(myTree, 30)

	sortSlice := []int{11, 5, 3, 30, 9, 20}
	Sort(sortSlice)
	fmt.Printf("Sorted values of the tree: %v\n", sortSlice)
}
