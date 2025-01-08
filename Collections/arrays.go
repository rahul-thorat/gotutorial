package collections

import "fmt"

func ArraysTest() {
	arr := [3]int32{1, 2}
	fmt.Println("Array:", arr)

	var arrDec [6]int64
	arrDec[3] = 3
	fmt.Println("Array:", arrDec)

	var arrDecB = [3]string{"1", "2"}
	fmt.Println("Array:", arrDecB)

	fmt.Println("ArrayLength:", len(arrDecB))

	for i := 0; i < len(arrDecB); i++ {
		fmt.Println(arrDecB[i])
	}

	for i, value := range arrDecB {
		fmt.Println(i, value)
	}

}

func SliceTest() {
	var slice []int32
	fmt.Printf("len: %v\n", len(slice))
	fmt.Printf("cap: %v\n", cap(slice))

	sliceV1 := []string{}
	fmt.Printf("len sliceV1: %v\n", len(sliceV1))
	fmt.Printf("cap sliceV1: %v\n", cap(sliceV1))

	sliceV1 = append(sliceV1, "hello")

	fmt.Println("SliceV1:", sliceV1)
	fmt.Printf("len sliceV1: %v\n", len(sliceV1))
	fmt.Printf("cap sliceV1: %v\n", cap(sliceV1))
}

func MapTest() {
	mapTest := map[int]string{}
	fmt.Println(mapTest)

	mapTestV1 := map[int]string{1: "apple", 2: "banana"}
	fmt.Println(mapTestV1)

	value, ok := mapTestV1[3]
	if ok {
		fmt.Println("Found orange:", value) // Output: Found apple: 5
	} else {
		fmt.Println("orange not found")
	}

	for key, value := range mapTestV1 {
		fmt.Println(key, value)
	}

}
