// Package integers provides basic arithmetic operations.
package integers

import "log"

// Add takes two integers and returns their sum.
//
// It also logs the result of the addition.
func Add(x, y int) (sum int) {
	sum = x + y
	log.Printf("sum of %d and %d is %d", x, y, sum)
	return
}
