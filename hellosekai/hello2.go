package main

import "fmt"

func main() {
	for j := 0; j < 150; j++ {
		if j % 2 == 0 {
			fmt.Print(j, ".")
		}
	}
}
