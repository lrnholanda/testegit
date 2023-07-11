package main

import "fmt"

func main() {
	x := 20

	func(x int) {
		fmt.Println(x, "vezes 3 é:")
		fmt.Println(x * 3)
	}(x)

	y := func(x int) int {
		return x * 2
	}
	fmt.Println(x, "vezes 2 é:", y(x))
}
