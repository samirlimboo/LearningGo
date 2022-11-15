package main

import (
	"fmt"
)

const aConst int = 62

func main() {
	var aString = "This is a ball"
	fmt.Println(aString)
	fmt.Printf("This is bussiness %T\n", aString)
	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("This is bussiness %T\n", anInteger)
	fmt.Println(aConst)
	fmt.Printf("this is a %T", aConst)
}
