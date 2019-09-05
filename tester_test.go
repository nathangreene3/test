package main

import (
	"fmt"
	"testing"
)

type tester struct {
	id     int
	values []int
}

func newTesterStruct(id int, value []int) tester {
	t := tester{id: id, values: make([]int, len(value))}
	copy(t.values, value)
	return t
}

func newTesterPointer(id int, values []int) *tester {
	t := &tester{id: id, values: make([]int, len(values))}
	copy(t.values, values)
	return t
}

func (t tester) compareStructs(u tester) int {
	switch {
	case t.id < u.id:
		return -1
	case u.id < t.id:
		return 1
	default:
		lent, lenu := len(t.values), len(u.values)
		min := lent
		if lenu < min {
			min = lenu
		}

		for i := 0; i < min; i++ {
			switch {
			case t.values[i] < u.values[i]:
				return -1
			case u.values[i] < t.values[i]:
				return 1
			}
		}

		switch {
		case lent < lenu:
			return -1
		case lenu < lent:
			return 1
		default:
			return 0
		}
	}
}

func (t *tester) comparePointers(u *tester) int {
	switch {
	case t.id < u.id:
		return -1
	case u.id < t.id:
		return 1
	default:
		lent, lenu := len(t.values), len(u.values)
		min := lent
		if lenu < min {
			min = lenu
		}

		for i := 0; i < min; i++ {
			switch {
			case t.values[i] < u.values[i]:
				return -1
			case u.values[i] < t.values[i]:
				return 1
			}
		}

		switch {
		case lent < lenu:
			return -1
		case lenu < lent:
			return 1
		default:
			return 0
		}
	}
}

func TestComparers(t *testing.T) {
	var (
		id    = 1
		value = []int{1, 2, 3}
		s, p  = newTesterStruct(id, value), newTesterPointer(id, value)
		r     int
	)

	if r = s.compareStructs(*p); r != 0 {
		t.Fatalf("expected 0\nreceived %d\n", r)
	}

	if r = s.comparePointers(p); r != 0 {
		t.Fatalf("expected 0\nreceived %d\n", r)
	}

	if r = (&s).comparePointers(p); r != 0 {
		t.Fatalf("expected 0\nreceived %d\n", r)
	}
}

func BenchmarkNewStruct(b0 *testing.B) {
	for i := 0; i < 10; i++ {
		values := make([]int, 0, i)
		for j := 0; j < i; j++ {
			values = append(values, j)
		}

		b0.Run(
			fmt.Sprintf("values = [0,%d]", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					newTesterStruct(i, values)
				}
			},
		)
	}
}

func BenchmarkNewPointer(b0 *testing.B) {
	for i := 0; i < 10; i++ {
		values := make([]int, 0, i)
		for j := 0; j < i; j++ {
			values = append(values, j)
		}

		b0.Run(
			fmt.Sprintf("values = [0,%d]", i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					newTesterPointer(i, values)
				}
			},
		)
	}
}

func BenchmarkSliceStructs(b0 *testing.B) {
	var (
		maxNumTesters = 10
		numValues     = 10
	)

	for i := 0; i < maxNumTesters; i++ {
		b0.Run(
			fmt.Sprintf("%d testers, %d values", i, numValues),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					s := make([]tester, 0, i)
					for k := 0; k < i; k++ {
						s = append(s, newTesterStruct(k, make([]int, numValues)))
					}
				}
			},
		)
	}
}

func BenchmarkSlicePointers(b0 *testing.B) {
	var (
		maxNumTesters = 10
		numValues     = 10
	)

	for i := 0; i < maxNumTesters; i++ {
		b0.Run(
			fmt.Sprintf("%d testers, %d values", i, numValues),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					s := make([]*tester, 0, i)
					for k := 0; k < i; k++ {
						s = append(s, newTesterPointer(k, make([]int, numValues)))
					}
				}
			},
		)
	}
}

func BenchmarkSliceStructs1(b0 *testing.B) {
	var (
		numTesters   = 10
		maxNumValues = 10
	)

	for i := 0; i < maxNumValues; i++ {
		b0.Run(
			fmt.Sprintf("%d testers, %d values", numTesters, i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					s := make([]tester, 0, numTesters)
					for k := 0; k < numTesters; k++ {
						s = append(s, newTesterStruct(k, make([]int, i)))
					}
				}
			},
		)
	}
}

func BenchmarkSlicePointers1(b0 *testing.B) {
	var (
		numTesters   = 10
		maxNumValues = 10
	)

	for i := 0; i < maxNumValues; i++ {
		b0.Run(
			fmt.Sprintf("%d testers, %d values", numTesters, i),
			func(b1 *testing.B) {
				for j := 0; j < b1.N; j++ {
					s := make([]*tester, 0, numTesters)
					for k := 0; k < numTesters; k++ {
						s = append(s, newTesterPointer(k, make([]int, i)))
					}
				}
			},
		)
	}
}

func BenchmarkSliceStructs2(b0 *testing.B) {
	var (
		maxNumTesters = 10
		maxNumValues  = 10
	)

	for h := 0; h < maxNumTesters; h++ {
		for i := 0; i < maxNumValues; i++ {
			b0.Run(
				fmt.Sprintf("%d testers, %d values", h, i),
				func(b1 *testing.B) {
					for j := 0; j < b1.N; j++ {
						s := make([]tester, 0, h)
						for k := 0; k < h; k++ {
							s = append(s, newTesterStruct(k, make([]int, i)))
						}
					}
				},
			)
		}
	}
}

func BenchmarkSlicePointers2(b0 *testing.B) {
	var (
		maxNumTesters = 10
		maxNumValues  = 10
	)

	for h := 0; h < maxNumTesters; h++ {
		for i := 0; i < maxNumValues; i++ {
			b0.Run(
				fmt.Sprintf("%d testers, %d values", h, i),
				func(b1 *testing.B) {
					for j := 0; j < b1.N; j++ {
						s := make([]tester, 0, h)
						for k := 0; k < h; k++ {
							s = append(s, newTesterStruct(k, make([]int, i)))
						}
					}
				},
			)
		}
	}
}
