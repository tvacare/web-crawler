package main

import "sort"

// paperSorte joins a By function and a slice of Papers to be sorted.
type paperSorter struct {
	papers []Paper
	by     func(p1, p2 *Paper) bool
}

// By is the type of a "less" function that defines the ordering of its Paper arguments.
type By func(p1, p2 *Paper) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(papers []Paper) {
	ps := &paperSorter{
		papers: papers,
		by:     by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *paperSorter) Len() int {
	return len(s.papers)
}

// Swap is part of sort.Interface.
func (s *paperSorter) Swap(i, j int) {
	s.papers[i], s.papers[j] = s.papers[j], s.papers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *paperSorter) Less(i, j int) bool {
	return s.by(&s.papers[i], &s.papers[j])
}
