// Go supports
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic factorial example.

package main

import (
	"fmt"
)

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fib(n int64) int64 {

	if n < 0 {
		if (n+1)%2 == 0 {
			return 1 * (fib((-1*n)-1) + fib((-1*n)-2))
		} else {
			return -1 * (fib((-1*n)-1) + fib((-1*n)-2))
		}

	}
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

func main() {
	//	fmt.Println(fib(-5))
	/*fmt.Println(fib(-6))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
		fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(-1))
	fmt.Println(fib(-2))
	*/

	fmt.Println(fib(-50))

}
