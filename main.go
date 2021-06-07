package main

import (
	"fmt"
	"strings"

	"github.com/FellPrimus/learngo/accounts"
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

	account := accounts.NewAccount("youngmin")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
