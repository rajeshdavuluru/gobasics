package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	
	fmt.Println(Vertex{1, 2})
	v := Vertex{3,4}
	
	v.X = 5
	fmt.Println(v);
	
	
	// Struct with pointers
	v1 := Vertex{1, 2}
	p := &v1
	p.X = 1e2
	fmt.Println(v1)
	
	
	v2 := Vertex{}
	fmt.Println(v2);
	
	v3 := Vertex{X: 10}
	fmt.Println(v3);
	
	v4 := Vertex{X: 10, Y:11}
	fmt.Println(v4);
	
}
