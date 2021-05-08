package main

import "fmt"

func main() {
	n := 2
	switch n {
	case 3:
		n = n-1
		fallthrough
	case 2:
		break
	case 1:
		n = 10
	}
	fmt.Println(n)
}
