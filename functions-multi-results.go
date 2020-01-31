package main

import "fmt"

// Function returning 2 results
func swap(x, y string) (string, string) {
	return y, x
}

func details() (int, string) {

	x:= 10 
	y:= "Rajesh"

	return x, y  
}

func main() {
	// returns two results same type (String)
	a, b := swap("R", "D")
	fmt.Println(a)
	fmt.Println(b)
	
	fmt.Println(b, "K", a);
	
	
	// returns two results different types
	var x int
	var y  string
	x, y = details()
	
	fmt.Println(x)
	fmt.Println(y)
}