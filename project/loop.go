package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
		if i<5 {
		fmt.Println("Ой!")
		}else {
		fmt.Println("Ай")
		}
    }
}