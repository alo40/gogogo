package main

import "fmt"

func main() {
	var myString = []rune("Résumé")
	var myRune = 'é'
	var indexed = myString[1]
	fmt.Println(myString[1])
	fmt.Printf("%v\n", indexed)
	fmt.Println(myRune)
	for i, v := range myString {
		fmt.Printf("%v %v %b\n", i, v, v)
	}
	fmt.Printf("The byte length of my string is %v\n", len(myString))
}
