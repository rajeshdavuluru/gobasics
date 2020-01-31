package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {


	// int variables without type
	x := 42
	y := 13
	fmt.Println(add(x, y))
	
	//int variables without out values
	z :=0
	fmt.Println(z);
	
	var i, j int = 1, 2
	
	fmt.Println(i)
	fmt.Println(j)
	
	//Variables with type
	var a int = 12
	var b int = 21
	
	fmt.Println(add(a, b))
	
	// int without 
	
	var c int 
	fmt.Println(c)  // Default value for int is 0
		
	//------------------------------------//
	//String variables with value
	var s string = "this is string"
	fmt.Println(s);
	
	
	//String variables without value
	var s1 string 
	fmt.Println(s1);  // Default value for string is EMPTY
	//------------------------------------//
	
	b1 := true
	var b2 bool
	
	fmt.Println(b1);
	fmt.Println(b2);
	
	
	
}