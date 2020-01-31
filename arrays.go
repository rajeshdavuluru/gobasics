package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	
	b := [3] string{"DRK", "DV", "DB"}
	fmt.Println(b[0]);

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	
	x :=[3] int {}
	fmt.Println(x) 
	
}