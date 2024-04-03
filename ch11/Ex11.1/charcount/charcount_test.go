package charcount

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	counts := map[rune]int{'a': 1, 'b': 1, 'c': 1}
	utflen := [utf8.UTFMax + 1]int{0, 3, 0, 0, 0}
	invalid := 0
	inputStr := "abc"
	readerInput := strings.NewReader(inputStr)

	c, u, i := CharCount(readerInput)

	for k, v := range c {
		if counts[k] != v {
			t.Errorf("CharCount(%q): invalid counts pair (%s,%d)", inputStr, string(k), v)
		}
	}

	if u != utflen {
		t.Errorf("CharCount(%q): invalid utflen counts %v", inputStr, u)
	}

	if i != invalid {
		t.Errorf("CharCount(%q): invalid = %d", inputStr, i)
	}
}
