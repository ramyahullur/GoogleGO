package main

import "fmt"

func main() {
	i := 31
	if i < 20 {
		fmt.Println("WOW You are young!! ")
	} else if i < 30 {
		fmt.Println("You are rocking your twenties!!")
	} else {
		fmt.Println("You are getting Old")
	}
}
