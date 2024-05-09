// https://www.acmicpc.net/problem/10872
// math, factorial

package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}
