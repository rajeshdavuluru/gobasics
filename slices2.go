package main

import "fmt"

func main() {
	a := make([]int, 5)  // creates a slice with length 5
	printSlice("a", a)

	b := make([]int, 2, 5) // creates a slice with length 0 and capacity of the underlaying array 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	
	
	var s [] int 
	fmt.Println(s , len(s), cap(s))
	
	s = append(s, 0)
	fmt.Println(s, len(s), cap(s))
	
	s = append(s, 1, 2, 3)
	fmt.Println(s, len(s), cap(s))
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
		
}
