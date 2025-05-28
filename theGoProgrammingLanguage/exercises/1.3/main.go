// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Echo1:", s, "Time:", time.Since(start).Nanoseconds(), "ns")
}

func echo2() {
	start := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println("Echo2:", s2, "Time:", time.Since(start).Nanoseconds(), "ns")
}

func echo3() {
	start := time.Now()
	fmt.Println("Echo3:", strings.Join(os.Args[1:], " "), "Time:", time.Since(start).Nanoseconds(), "ns")
}

func main() {
	echo1()
	echo2()
	echo3()
}
