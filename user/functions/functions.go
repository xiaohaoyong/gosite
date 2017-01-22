package main

import (
	"fmt"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// const Pi = 3.14

// func main() {
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")
// 	const Truth = true
// 	fmt.Println("G rules?", Truth)
// }

// func main() {
// 	v := 42
// 	fmt.Printf("v is of type %Tn", v)
// }

// func main() {
// 	var x, y int = 3, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z int = int(f)
// 	fmt.Println(x, y, z)
// }

// var (
// 	Tobe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	var f float64
// 	fmt.Println(f)
// }

// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no!"
// 	fmt.Println(i, j, k, c, python, java)
// }

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }
// func main() {
// 	fmt.Println(split(17))
// }
