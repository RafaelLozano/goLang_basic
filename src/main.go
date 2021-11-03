package main

import (
	"fmt"
)

//practice reverse a string this is a helper function
func reverseString(s string) string{
	var reverse string
	for i := len(s) - 1; i>= 0; i--{
		reverse += string(s[i])
	}
	return reverse
}
func isPalindrome(s string) {
	if s == reverseString(s){
		fmt.Printf("%s is palindrome", s)
	}else{
		fmt.Printf("%s is not palindrome",s)
	}
}

func main() {
	isPalindrome("racecar")
}