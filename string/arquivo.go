package string

import "fmt"

func init() {
	s := "hello, world"

	fmt.Println(string(s[0]), s[7])

	// s[0] = "g"
}