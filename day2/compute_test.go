package main

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	cases := []struct {
		Input  []int
		Output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, o := range cases {
		p := &Program{Tape: o.Input, CurrentPosition: 0}
		p.Execute()
		if !reflect.DeepEqual(p.Tape, o.Output) {
			t.Errorf("Expected: %v, Got: %v", o.Output, p.Tape)
		}
	}
}
