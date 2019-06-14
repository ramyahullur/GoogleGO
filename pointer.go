package main

import (
	"fmt"
	"time"
)

func max(i int, j int, k *int) {
	fmt.Println("the memory location of k in the stack is")
	fmt.Println(&k)
	if i > j {
		*k = i
	} else {
		*k = j
	}
}
func main() {
	var c int
	fmt.Println("the memory location of c is")
	fmt.Println(&c)
	max(100, 345, &c)
	fmt.Println("the highest number is")
	fmt.Println(c)
	time.Sleep(time.Second)
	fmt.Println("the memory location of c in the main function has been reffered to the memory location of k in the max function")
	fmt.Println("which is in the stack memory and is allocated to the values of i or j depending on the condition they meet")

}
