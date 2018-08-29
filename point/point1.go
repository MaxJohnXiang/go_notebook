package main

import "fmt"

func main() {
	x := 1
	p := &x
	// 那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是*int，
	fmt.Println(*p)
	*p = 2
	fmt.Println("x is", x)
	fmt.Println()
	fmt.Println(f() == f())
	fmt.Println("-------------")
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
