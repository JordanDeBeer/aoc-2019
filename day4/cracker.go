package main

import (
	"strconv"
	"strings"
)

func FindPasswords(min, max int) []int {
	var passwords []int
	for v := range GenerateNumbers(min, max) {
		passwords = append(passwords, v)
	}
	return passwords
}

func GenerateNumbers(min, max int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for p := min; p <= max; p++ {
			if CheckDoubles(p) && CheckNotDecreasing(p) && CheckIsolatedDoubles(p) {
				ch <- p
			}
		}
	}()
	return ch
}

func CheckDoubles(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}
func CheckIsolatedDoubles(n int) bool {
	s := strconv.Itoa(n)
	var hasMatch bool
	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			if strings.Count(s, string(s[i])) == 2 {
				hasMatch = true
			}
		}
	}
	return hasMatch
}

func CheckNotDecreasing(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < 5; i++ {
		s1, _ := strconv.Atoi(string(s[i]))
		s2, _ := strconv.Atoi(string(s[i+1]))
		if s1 > s2 {
			return false
		}
	}
	return true
}
