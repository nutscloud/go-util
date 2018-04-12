package slice

import (
	"testing"
)

type MatchInSliceType struct {
	sli      []interface{}
	elems    interface{}
	expected bool
}

func TestMatchInSlice(t *testing.T) {
	matchs := []MatchInSliceType{
		MatchInSliceType{
			sli:      []string{"aaa", "ccc"},
			elems:    "aaa",
			expected: true,
		},
		MatchInSliceType{
			sli:      []string{"aaa", "ccc"},
			elems:    "aaaf",
			expected: false,
		},
		MatchInSliceType{
			sli:      []int{1, 2, 3, 10},
			elems:    5,
			expected: false,
		},
		MatchInSliceType{
			sli:      []int{1, 2, 3, 10},
			elems:    2,
			expected: true,
		},
	}

	for _, m := range matchs {
		if m.expected != MatchInSlice(m.sli, m.elems) {
			t.Fatalf("Failed to test with '%s': %s, '%v'", test.fileName,
				test.name, realError)
		}
	}
}
