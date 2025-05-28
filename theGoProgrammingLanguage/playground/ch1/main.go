package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// echo1
	fmt.Println("Hello" + "Hi") // includes a new line at the end
	fmt.Println("Second line")

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " // During first iteration sep is empty ("") therefore no space at the beginning.
	}
	fmt.Println(s)

	//for loop GO, this is the only loop statement in GO

	// for initialization; condition; post {
	// 	 Zero or more statements
	// }

	// a traditional while loop
	// for condition {
	// 	// ...
	// }

	// // infinite loop
	// for {
	// 	// ...
	// }

	s2, sep2 := "", ""                //echo2
	for _, arg := range os.Args[1:] { // Range produces a pair of values
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)

	fmt.Println(strings.Join(os.Args[1:], " ")) // echo3 = most efficient way

	fmt.Println(os.Args[1:]) // if we just want to see the values and dont care about formatting, any slice can be printed this way

}
