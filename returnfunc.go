package main

import "fmt"

func max(i int, j int) {
	fmt.Scanln(&i, &j)
	if i > j {
		fmt.Println(i)
	} else {
		fmt.Println(j)
	}
}

func main() {
	fmt.Println("Enter two intergers that need to be compared")
	max(2, 3)
}
