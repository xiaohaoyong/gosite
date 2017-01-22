package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X)
// 	fmt.Println(Vertex{1, 2})
// }

// func main() {
// 	i, j := 42, 2701
// 	p := &i
// 	fmt.Println(*p)
// 	*p = 21
// 	fmt.Println(i)

// 	p = &j
// 	*p = *p
// 	fmt.Println(j)
// }
