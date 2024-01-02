package problem0013

import "fmt"

func RomanToInt(s string) int {

	sum := 0

	intForRoman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	fmt.Println(intForRoman)
	return sum
}
