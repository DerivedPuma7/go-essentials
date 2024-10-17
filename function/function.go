package function

import "fmt"

func init() {
	funcAnonima()
}

func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func funcAnonima() {
	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	
	fmt.Printf("%T \n", squares)
	fmt.Printf("%T \n", f)
}