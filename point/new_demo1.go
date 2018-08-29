package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	a := new(int)
	b := new(int)
	fmt.Println(*a, *b)
	fmt.Println(a == b)
}
