package slice

import "fmt"

func init() {
	// testeSlice()
	testeSlice2()
}

func testeSlice() {
	months := [...] string { 1: "jan", 2: "fev", 3: "mar", 4: "abril", 5: "maio", 6: "jun", 7: "jul", 8: "ago", 9: "set", 10: "out", 11: "nov", 12: "dez" }
	
	firstSemester1 := months[1:7]
	firstSemester2 := months[:7]
	fmt.Println(firstSemester1)
	fmt.Println(firstSemester2)
	
	secondSemester1 := months[7:13]
	secondSemester2 := months[7:]
	fmt.Println(secondSemester1)
	fmt.Println(secondSemester2)
}

func testeSlice2() {
	months := [] string { 1: "jan", 2: "fev", 3: "mar", 4: "abril", 5: "maio", 6: "jun", 7: "jul", 8: "ago", 9: "set", 10: "out", 11: "nov", 12: "dez" }
	fmt.Printf(fmt.Sprintf("%q", months))
}