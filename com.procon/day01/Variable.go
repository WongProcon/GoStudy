package main

import "fmt"

//全局变量m
var m int32 = 100

const (
	n1 = iota
	n2
	n3
	n4
)

// 匿名变量用一个下划线表示
func foo() (int, string) {
	return 10, "procon"
}

func main() {
	//局部变量n
	n := 300
	println(m, "\t", n)
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println(n1, n2, n3, n4)
}
