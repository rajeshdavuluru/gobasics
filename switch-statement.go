package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	
	
	// Switch with no condiftion
	x := 12
	switch {
	case x % 2 == 0:
		fmt.Println("Positive Even")
	case x % 2 > 0:
		fmt.Println("Positive Odd")
	default:
		fmt.Println("-ve Number")
	}
}