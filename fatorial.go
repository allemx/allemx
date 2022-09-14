package main

import "fmt"

func main() {
	for i := 21; i> 0;i--{
	res := Fatorial(i)
	fmt.Println("Res : ", i , res)}
	
}
func Fatorial(x int) int {
	var y  = 1
	for ; x > 0; x-- {
		y = y * x
	}
	return y
}
