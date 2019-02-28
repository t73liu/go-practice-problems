package math

import "strconv"

func fizzBuzz(n int) []string {
	arr := make([]string, n)
	for i := 1; i <= n; i++ {
		mod3 := i%3 == 0
		mod5 := i%5 == 0
		if mod3 && mod5 {
			arr[i-1] = "FizzBuzz"
		} else if mod3 {
			arr[i-1] = "Fizz"
		} else if mod5 {
			arr[i-1] = "Buzz"
		} else {
			arr[i-1] = strconv.Itoa(i)
		}
	}
	return arr
}
