package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming", "algorithms"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	var stack []string
	stackMap := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	pop := func(s map[string]bool, key string) (bool, bool) {
		v, ok := s[key]
		if ok {
			delete(s, key)
		}
		return v, ok
	}
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				stack = append(stack, item)
				stackMap[item] = true
				visitAll(m[item])
				order = append(order, item)
				pop(stackMap, item)
				stack = stack[:len(stack)-1]
				continue
			}
			if _, ok := stackMap[item]; ok {
				fmt.Printf("Cycle found:\n")
				for i, course := range stack {
					if i == len(stack)-1 {
						fmt.Printf("%s", course)
						break
					}
					fmt.Printf("%s <- ", course)
				}
				fmt.Printf("\n")
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
