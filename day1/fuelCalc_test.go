package main

import (
	"testing"
)

func TestFuelCalc(t *testing.T) {
	cases := []struct {
		Input  int
		Output int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, o := range cases {
		req := FuelCalc(o.Input)
		if req != o.Output {
			t.Errorf("Expected: %v, Got: %v", o.Output, req)
		}
	}
}
