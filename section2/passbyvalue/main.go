package main

import "fmt"

func main() {
	x := 5
	wontchangevalue(x)
	fmt.Println(x)
	changevalue(&x)
	fmt.Println(x)
}

func wontchangevalue(val int) {
	val = 15
}

func changevalue(val *int) {
	*val = 10
}
