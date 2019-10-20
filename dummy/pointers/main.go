package main

import "fmt"

func main() {
	i, j := 100, 500
	p := &i
	fmt.Println(*p)
	*p = 200
	fmt.Println(i)
	p = &j
	*p = *p / 2
	fmt.Println(j)
}