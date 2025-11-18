package main

import ( 
	"fmt"
	"math"
)

func main() {

	var a = "word1"
	var b, c int = 1, 2
	var d int

	fmt.Println("Hello, 世界")
	fmt.Println("2 x 4 = ", 2 * 4)
	fmt.Println(true || false)
	fmt.Println(a)
	fmt.Println(b + c)
	fmt.Println(d)

	e := "apple"
	fmt.Println(e)
	
	const f = 50000
	fmt.Println(math.Sin(f))
}
