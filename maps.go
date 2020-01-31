package main

import "fmt"


var nilmap map[string]int


func main(){

	fmt.Println(nilmap);
	//nilmap["Zero"] = 0  // we cannot assign a value to nil map
	
	
	nilmap = make(map[string]int)
	fmt.Println(nilmap);
	nilmap["Zero"] = 0   // after make a map only we can assign a value
	fmt.Println(nilmap);

	var evenmap = make(map[string]int)
	
	evenmap["Two"] = 2
	evenmap["Four"] = 4
	evenmap["Six"] = 6
	evenmap["Eight"] = 8
	evenmap["Ten"] = 10

	fmt.Println(evenmap);
	
	
	var oddmap = map[string]int{
		"One" :1,
		"Three": 3,
		"Five" : 5,
		"Seven" : 7,
		"Nine" : 9,
	}
	
	fmt.Println(oddmap)
	
	
	var moviemap = make(map[string]string)

	moviemap["favoriteMovie"] = "ICE AGE"
	
	fmt.Println(moviemap)
	
	// updating 
	moviemap["favoriteMovie"] = "ICE AGE 2"
	fmt.Println(moviemap)
	
	//delete 
	delete(moviemap, "favoriteMovie")
	fmt.Println(moviemap)
	
	
	value, status := moviemap["favoriteMovie"]
	fmt.Println("The value:", value, "Present?", status)
}