package main

import "fmt"

func main() {
	inc := incrementor()
	i := 0
	for i < 100 {
		fmt.Println(inc())
		i++
	}
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
