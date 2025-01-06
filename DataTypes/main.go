package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 -> unsigned
	var num int64 = 12345
	fmt.Println(num)

	//float32 float64
	floatNo := 234.23
	fmt.Println(floatNo)

	//perform operations on same datatypes only or use cast
	//addition := float64(num) + floatNo -> compile time error
	addition := float64(num) + floatNo
	fmt.Println(addition)

	//string
	var myStr string = "Hello ε World"
	fmt.Println(myStr)

	//no of bytes in string
	fmt.Println(len(myStr))
	//length of string
	fmt.Println(utf8.RuneCountInString(myStr))

	var myBool bool = true
	fmt.Println(myBool)

	const pi float64 = 3.14
	//pi = 3.23 -> compile error cannot change const

	var1, var2, var3 := 1, "Hello", false
	fmt.Printf("%d %s %t\n", var1, var2, var3)

}
