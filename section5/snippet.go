package main

import "fmt"

// Function to calculate the factorial of a number
// This function doesn't return the factorial

// func factorial(n int) int {
//	The expected result is an int and not a float
// 	result := 1.00
// 	for i := 1; i <= n; i++ {
// 		// Here because of Go being a strictly typed language 
// 		// we cannot reassign an integer to a memory that previous held a float
// 		// neither can we carry out mathematical operations between 2 different types
		// Also it should be multiplication and not addition
// 		result += i
// 	}
// 	return result
// }


// Possible Improvement to the fixed code would be to take into consideration that
// n might become very large, standard integer types might overflow 
// Hence using a library that supports big integers (like math/big in Go) to ensure correctness for large values.
// func factorial(n int64) *big.Int {
//     result := big.NewInt(1)
//     for i := int64(1); i <= n; i++ {
//         result.Mul(result, big.NewInt(i))
//     }
//     return result
// }

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Function to print the factorial of a number
func printFactorial() {
	num := 5
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}


func main() {
	// Here the printfactorial isn't referenced to any declared function as Go is case sensitive
	// printfactorial()
	printFactorial()
}
