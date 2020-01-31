package main

import "fmt"

// No Braces () required for loop , but parenthesis {} is required
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	// No init and post statements, no colons also.
	// this will be useful as a while
	sum = 0
	i := 0
	for i < 10{
		sum += i
		i++ 
	}
	fmt.Println(sum)
	
	// No init and post statements, with colons 
	sum = 1
	for ; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
	
	//infinite loop
	for {
	
	}
	
}
