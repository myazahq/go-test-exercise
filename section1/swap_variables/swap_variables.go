package swapvariables

import (
	"errors"
	"reflect"
	"fmt"
)

// Take the addresses of the variable
// Compare if the addresses contain variables of the the same type
// if true --> then proceed to swap
	// - if the variable type isn't available return unsupported type error
// else return an error saying swap is impossible due to the different types

func SwapVariables(a, b interface{}) error {
	if areVariablesOfTheSameType(a, b){
		switch a.(type) {
		case *int:
			*a.(*int), *b.(*int) = *b.(*int), *a.(*int)
		case *string:
			*a.(*string), *b.(*string) = *b.(*string), *a.(*string)	
		case *float64:
			*a.(*float64), *b.(*float64) = *b.(*float64), *a.(*float64)
		// We can add more cases here to work for other types
		default:
			return fmt.Errorf("unsupported type: %T", a)
		}

		return nil
	}
	return errors.New("variables are of different types, therefore cannot be swapped")
}

func areVariablesOfTheSameType(a, b interface{}) bool {
    type1 := reflect.TypeOf(a)
    type2 := reflect.TypeOf(b)
    return type1 == type2
} 