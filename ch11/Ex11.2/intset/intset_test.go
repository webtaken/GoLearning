package intset

import "testing"

func TestEmptyIntSet(t *testing.T) {
	emptySet := IntSet{}
	addedValue := 4
	if emptySet.Has(addedValue) == true {
		t.Errorf("Has(%q) = true", addedValue)
	}
}

func TestIntSetAdd(t *testing.T) {
	set := IntSet{}
	values := []int{1, 2, 3, 4, 78, 45}
	for _, v := range values {
		set.Add(v)
	}
	for _, v := range values {
		if set.Has(v) != true {
			t.Errorf("Has(%q) = false", v)
		}
	}
}

func TestIntSetAddNegativeValues(t *testing.T) {
	set := IntSet{}
	values := []int{1, -2, 3, -4, 78, -45}
	for _, v := range values {
		set.Add(v)
	}
	for _, v := range values {
		if v < 0 && set.Has(v) == true {
			t.Errorf("Has(%q) = true", v)
		}
		if v >= 0 && set.Has(v) != true {
			t.Errorf("Has(%q) = false", v)
		}
	}
}
