package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": false},
	"calculus":   {"linear algebra": false},
	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},
	"data structures":       {"discrete math": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	order := []string{}
	seen := make(map[string]bool)
	var visitAll func(course string, items map[string]bool)
	visitAll = func(course string, items map[string]bool) {
		for key1 := range items {
			if !seen[key1] {
				seen[key1] = true
				visitAll(key1, m[key1])
				order = append(order, key1)
			}
		}
		if !seen[course] {
			seen[course] = true
			order = append(order, course)
		}
	}
	for course, prereqs := range m {
		visitAll(course, prereqs)
	}
	return order
}
