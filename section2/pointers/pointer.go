package main

import "fmt"

func main() {
	var p1 *int
	d1 := 4
	fmt.Println(&d1)
	p1 = &d1
	fmt.Println(p1)
	fmt.Println(*p1)
}
