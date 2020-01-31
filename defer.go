package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	
	// defer statements inserted into stack and retrieved LIFO order
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
