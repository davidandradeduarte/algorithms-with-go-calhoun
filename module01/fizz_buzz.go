package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var result strings.Builder

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result.WriteString("Fizz Buzz, ")
		} else if i%3 == 0 {
			result.WriteString("Fizz, ")
		} else if i%5 == 0 {
			result.WriteString("Buzz, ")
		} else {
			result.WriteString(fmt.Sprintf("%d, ", i))
		}
	}

	// we could refactor this by adding:
	/*
		if i != n {
			result.WriteString(", ")
		}
	*/
	// inside the for loop
	if result.Len() > 2 && result.String()[result.Len()-2:] == ", " {
		tmp := result.String()
		result.Reset()
		result.WriteString(tmp[:len(tmp)-2])
	}

	fmt.Println(result.String())
}
