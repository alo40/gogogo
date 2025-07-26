package main

import "fmt"

func main() {
	var intArr = [...]int32{0, 1, 2, 3, 4} // size of the array infered by ...
	// intArr[0] = 1
	// intArr[1] = 2
	// intArr[2] = 3
	// intArr[3] = 4
	// intArr[4] = 5
	fmt.Println(intArr)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	fmt.Println(&intArr[3])
	fmt.Println(&intArr[4])
}
