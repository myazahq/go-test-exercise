package main

import (
	"fmt"

	"github.com/myazahq/go-test-exercise/section1/log"
	"github.com/myazahq/go-test-exercise/section1/slices"
	swapvariables "github.com/myazahq/go-test-exercise/section1/swap_variables"
)


func main(){

	// Section 1
	fmt.Println("**** SWAP VARIABLES EXERCISE ****")
	variableA := 8
	variableB := 9
	err := swapvariables.SwapVariables(&variableA, &variableB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(variableA, variableB) //expects variableA = 9 and variableB = 8

	fmt.Println("**** SUM OF EVEN NUMBERS EXERCISE ****")
	fmt.Printf("expected: %d got: %d", 12, slices.SumOfEvenNumbers([]int{1,2,3,4,6}))

	fmt.Println("**** Logger Interface ****")
	fileLogger := log.NewFileLogger("log.txt")
	fileLogger.Log("Hello, Logging from File Logger\n")

	consoleLogger := log.ConsoleLogger{}
	consoleLogger.Log("Hello, Logging from Console Logger")

}