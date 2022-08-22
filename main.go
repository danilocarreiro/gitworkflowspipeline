package main

import (
	"fmt"
	"gwfp/src"
)

func main() {
	avg := src.ToFixed(src.Average(8, 10, 5.6, 3, 2.89), 2)
	fmt.Println(avg)
	avg = src.ToFixed(src.Average(8, 2, 3, 5), 2)
	fmt.Println(avg)
}
