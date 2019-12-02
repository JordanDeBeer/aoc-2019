package main

func FuelCalc(n int) int {
	total := (n / 3) - 2
	if total > 0 {
		total += FuelCalc(total)
	} else {
		return 0
	}
	return total
}
