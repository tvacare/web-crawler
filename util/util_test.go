package util_test

import (
	"testing"

	"github.com/tvacare/web-crawler/util"
)

func TestSliceContains(t *testing.T) {
	s1 := "pineapple"
	slice1 := []string{"apple", "banana", "orange", "pear"}
	b1, p1 := util.SliceContains(s1, slice1)

	if b1 == false || p1 == "" {
		t.Errorf("Slice contains should have matched - %s, %v", s1, slice1)
	}

	s2 := "car"
	slice2 := []string{"bus", "subway", "airplane", "train"}
	b2, p2 := util.SliceContains(s2, slice2)

	if b2 != false || p2 != "" {
		t.Errorf("Slice contains should not match - %s, %v", s2, slice2)
	}
}
