package go_pkg

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("%d %d", len(s), cap(s))
}
