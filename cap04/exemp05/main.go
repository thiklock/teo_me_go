package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")

	default:
		fmt.Println("Unknown")
	}
}
