package main

import (
	"fmt"
	"time"
)

func isfloat64(in interface{}) bool {

	if _, ok := in.(float64); ok {
		fmt.Println(in)
		return true
	} else {
		return false
	}
}

// func area[t int | any](a, b t) t {

// 	ar := a * b
// 	return ar
// }

func main() {
	start := time.Now().Nanosecond()
	// fmt.Println(area(3, 4))
	// fmt.Println(area(2.34243, 2.344))
	fmt.Println(isfloat64(34))
	fmt.Println(isfloat64(34.432))
	end := time.Now().Nanosecond()
	// fmt.Println(end.Sub(start))
	fmt.Println(end - start)
}
