package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening")
// 	}
// }

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	fmt.Println(today)
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today")
// 	case today + 1:
// 		fmt.Println("Tomorrow")
// 	case today + 2:
// 		fmt.Println("In two days")
// 	default:
// 		fmt.Println("Too far away")
// 	}
// }

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		fmt.Println("linux .")
// 	default:
// 		fmt.Printf("%s", os)
// 	}
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	}
// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

// func main() {
// 	fmt.Println(sqrt(2), sqrt(-4))
// }

// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 		fmt.Println(i)
// 	}
// }
