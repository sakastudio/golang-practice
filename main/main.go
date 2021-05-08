package main

import "fmt"

func main() {
	a,b := 1000,100
	if a>b{
		fmt.Println("aのほうがデカい")
	}else if a==b{
		fmt.Println("aとb同じ")
	}else {
		fmt.Println("bのほうがデカい")
	}
}
