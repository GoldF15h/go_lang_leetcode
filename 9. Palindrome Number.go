package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
    s := []rune(strconv.FormatInt(int64(x),10))
	var rev_s []rune
	// fmt.Println("STR :",s)
	for i := len(s)-1 ; i >= 0 ; i-- {
		// fmt.Println(string(s[i]))
		rev_s = append(rev_s, s[i])
	}
	
	for i := 0 ; i < len(s) ; i++{
		if s[i] != rev_s[i] {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println("test")
	var inp int
	// fmt.Scanln(&inp)
	inp = 121
	fmt.Println(isPalindrome(inp))
}

// func main() {
//   i := 10
//   s1 := strconv.FormatInt(int64(i), 10)
//   s2 := strconv.Itoa(i)
//   fmt.Printf("%v, %v\n", s1, s2)
// }