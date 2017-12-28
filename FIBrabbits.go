
package main

import (
	"fmt"
)

func fibonacci(nummonths int, numkids int) int{

	output := 0

	if nummonths == 1 || nummonths == 0 {
		output = nummonths

	}else if nummonths == 2 {
		output = numkids

	}else if genone, gentwo, factor := (fibonacci((nummonths) - 1, numkids)),
					   (fibonacci((nummonths) - 2, numkids)),
					   1; nummonths > 2 {
		if nummonths > 4 {
			factor = numkids
		}
		output = (genone + (gentwo * factor))
	}
	return output
}

func main() {
	fib := fibonacci (29,2)
	fmt.Println(fib)
}
