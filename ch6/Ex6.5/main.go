package main

import (
	"bytes"
	"fmt"
)

const SIZE = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/SIZE, uint(x%SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/SIZE, uint(x%SIZE)
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var newIntSet *IntSet = &IntSet{words: make([]uint, len(s.words))}
	copy(newIntSet.words, s.words)
	return newIntSet
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		tmp := word
		for tmp != 0 {
			tmp = tmp & (tmp - 1)
			count++
		}
	}
	return count
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, SIZE*i+j)
			}
		}
	}
	return elems
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	fmt.Println("Machine arquitecture:", SIZE)

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // "3"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // "2"
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())    // "4"
	x.Remove(1)
	fmt.Println(x.String())           // "{9 42 144}"
	fmt.Println(x.Len())              // "3"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	y.Clear()
	fmt.Println(y.String()) // "{}"
	z := x.Copy()
	fmt.Println(z.String()) // "{9 42 144}"

	x.AddAll(5, 10, 3)
	fmt.Println(x.String()) // "{3 5 9 10 42 144}"

	A := x.Copy() // "{3 5 9 10 42 144}"
	A.IntersectWith(z)
	fmt.Println(A) // "{9 42 144}"

	z.DifferenceWith(A)
	fmt.Println(z) // "{}"

	B := x.Copy() // "{3 5 9 10 42 144}"
	B.SymmetricDifference(A)
	fmt.Println(B.String()) // "{3 5 10}"

	for _, elem := range x.Elems() {
		fmt.Printf("%v ", elem*elem) // "9 25 81 100 1764 20736"
	}
}
