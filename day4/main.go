package main

import "fmt"

func main() {

	min := 240298
	max := 784956
	passwords := FindPasswords(min, max)
	fmt.Printf("Passwords: %v. Length: %v\n", passwords, len(passwords))
}
