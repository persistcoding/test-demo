package main

import "fmt"

//get sum
func getSum(a int, b int) int {
	c := a + b
	return c
}

func main() {
	d := getSum(1, 2)
	fmt.Println(d)
}
