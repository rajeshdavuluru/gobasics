package main

import "fmt"

func main(){

	x := checkEvenOdd(10)
	fmt.Println(x);
	
	fmt.Println(checkEvenOdd(19));
	
	//if with short statements
	if y:=13; y%2 ==0{
		fmt.Println("Even");
	}else {
		fmt.Println("Odd");
	}
	
}

func checkEvenOdd(x int) string{

	var result string
	
	if x %2 ==  0{
		result =  "Even"
	}else{ 
		result =  "Odd"
	}
	
	return result
}
	