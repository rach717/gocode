package doer

import "fmt"

type Doer interface {
	CheckPalindrome(string) bool
}


func CheckPalindrome(input string) bool {

	fmt.Println("inputstring in function:", input)
	var result string
	for _, v := range input {
		result = string(v) + result
		fmt.Println(result)
	}
	return input == result
}
