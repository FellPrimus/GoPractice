package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("youngmin")
	fmt.Println(totalLength, upperName)
	names := []string{"y", "o", "u", "n", "g"}
	names = append(names, "m")
	fmt.Println(names)

	repeatMe("young", "min", "go", "study")
}
