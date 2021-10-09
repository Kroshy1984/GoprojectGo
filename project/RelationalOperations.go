package main

import "fmt"

func main() {
	x:=42
	y:=8
	fmt.Println(x==y)
	fmt.Println(x!=y)
	fmt.Println(x>y)
	fmt.Println(x<y)
	fmt.Println(x!=y && x<=y) // AND
	fmt.Println(x!=y || x<=y) // OR
	fmt.Println(!(x>y))
}