package main

import "fmt"

func main() {
	var array1 = [5]float64{0, 1, 2, 3, 4}
	fmt.Printf("The original array values are: %v\n", array1)
	fmt.Printf("Memory location of array1 is %p\n", &array1)
	// for index, value := range array1 {
	// 	fmt.Println(index, value, &value)
	// }
	var result [5]float64 = square(&array1)
	fmt.Printf("The square array values are: %v\n", result)
}

func square(array *[5]float64) [5]float64 {
	fmt.Printf("Memory location of array inside square function is %p\n", array)
	// fmt.Printf("Memloc of array pointer %p, memloc of array %p\n", array, &array)
	for i := range array {
		array[i] = array[i] * array[i]
	}
	return *array
}
