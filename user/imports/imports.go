package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
