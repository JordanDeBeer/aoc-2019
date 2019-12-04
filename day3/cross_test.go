package main

import (
	"testing"
)

func TestFindClosestCross(t *testing.T) {
	cases := []struct {
		InputOne []string
		InputTwo []string
		Output   int
	}{
		{[]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, 6},
		{[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, 159},
		{[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, 135},
	}
	for i, o := range cases {
		w1 := NewWire(o.InputOne)
		w2 := NewWire(o.InputTwo)
		x := FindCrosses(w1, w2)
		t.Logf("Crosses: %v", x)
		y := FindClosestCross(x)
		if y != o.Output {
			t.Errorf("Failed on Case: %v, Expected: %v, Got: %v", i, o.Output, y)
		}
	}
}
