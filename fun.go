package main

import "fmt"
import "time"

func main() {
	for day := 1; day < 4; day++ {
		fmt.Println("on the day", day)
		switch day {
		case 1:
			{
				fmt.Println("he saw me")
			}
			time.Sleep(time.Second)
		case 2:
			{
				fmt.Println("he liked me")
			}
			time.Sleep(time.Second)
		case 3:
			{
				fmt.Println("he proposed me!")
			}
		}
	}
	fmt.Println("We are in Love!! :) ")
}
