package main

import "fmt"

func main() {
	
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	
	//Array Slices
		
	fmt.Println(primes[1:4])
	fmt.Println(primes[1:])
	fmt.Println(primes[:6])
	fmt.Println(primes[:])
	fmt.Println(primes)
	//fmt.Println(primes[])  // this is not a valid statement
	
	
	//Slices dont store any data
	// if you change the value of the slice it will change the value in underlaying array
	
	even :=[10] int{2,4,6,8,10,12,14,16,18,20}
	
	x :=even[2:6]
	y := even[5:8]
	fmt.Println(x);
	fmt.Println(y)
	
	x[3] = 22
	fmt.Println(x);
	fmt.Println(y)
	fmt.Println(even)
	
	// Slice literals are same like array literals but no length
	
	z :=[] int {}   // Slice Literal
	fmt.Println(z)  // displays [] empty array
	
	z1 :=[3] int {}  //Array Literal
	fmt.Println(z1)  //displays [0,0,0]
	
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	
	// Slice Length and capacity
	odd := [5] int {1,3,5,7,9}
	
	oddSlice := odd[2: 4]
	// Length  =  Length of the slice
	//Capacity = size of the underlaying array starting from slice first element. in this case slice start at  5 , so capacity is 3
	fmt.Printf("Length %d : Capacity %d\n" , len(oddSlice), cap(oddSlice));  //Length 2 : Capacity 3
	
	
	// Nil Slices
	
	var nilslice []int
	fmt.Println(nilslice, len(nilslice), cap(nilslice))
	if nilslice == nil {
		fmt.Println("nil!")
	}
	
	
}