package main

import (
	"fmt"
	"strings"
)

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
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	names := [5]string{"y", "o", "u", "n", "g"}
	fmt.Println(names)

	repeatMe("young", "min", "go", "study")
}
