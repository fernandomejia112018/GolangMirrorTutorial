// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 5 values in the `queue` channel.
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	queue <- "five"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
