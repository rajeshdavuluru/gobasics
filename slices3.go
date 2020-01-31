package main

import "fmt"

// Range Function returns 2 parameters. index and value
func main(){

	var even = []int {2,4,6,8,10,12}
	
	for index, value := range even{
	
		fmt.Printf("index: %d - value: %d \n", index, value);
	}

	// if you dont want index you can use underscore(_)
	for _, value := range even{
	
		fmt.Printf("value: %d \n",  value);
	}
	
	// if you dont want value you can  do 2 ways 1) use underscore(_) 2) dont mention second variable
	for index, _ := range even{
	
		fmt.Printf("index: %d  \n", index);
	}
	
	for index := range even{
	
		fmt.Printf("index: %d  \n", index);
	}
}