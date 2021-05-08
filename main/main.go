package main

import "fmt"

func main() {
	x,y := 3,4
	x,y = swap(x,y)

	fmt.Println(x,y)

	x,_ = swap(x,y)

	fmt.Println(x)
}

func swap(i, j int) (int,int) {
	return j,i
}
