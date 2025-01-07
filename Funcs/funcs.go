package funcs

import (
	"errors"
	"fmt"
)

func FuncsCaller() {
	printVal("Hello World")
	var numerator int = 11
	denominator := 2
	result, remainder, err := integerDiv(numerator, denominator)
	switch {
	case (err != nil):
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of division is %v\n", result)
	default:
		fmt.Printf("The result of division is %v and remainder is %v\n", result, remainder)
	}

}

func printVal(printVal string) {
	if len(printVal) == 0 {
		fmt.Println("Empty String")
		return
	}
	fmt.Println(printVal)
}

func integerDiv(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err

}
