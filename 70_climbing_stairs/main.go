package main

import "fmt"

func climbStairs(n int) int {
	one := 1
	two := 1

	for i := 0; i < (n - 1); i++ {
		temp := one
		one = one + two
		two = temp
	}

	fmt.Println(one)
	return one

}

//TODO: research fibonacci

func main() {
	climbStairs(5)
}
