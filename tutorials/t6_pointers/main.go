package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	*p = 10
	fmt.Printf("The memory location the pointer p points is %v\n", p)
	fmt.Printf("The value the pointer p points is %v\n", *p) // Dereferencing the pointer
	p = &i                                                   // p refers to the memory address of i
	*p = 100
	fmt.Printf("The new memory location the pointer p points is %v\n", p)
	fmt.Printf("The new value the pointer p points is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)
}
