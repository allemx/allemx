package main

import "fmt"

func main() {
	var i float64
	for i = 21; i > 0; i-- {
		res := Fatorial(i)
		fmt.Println("Res : ", i, res)
	}

}
func Fatorial(x float64) float64 {
	var y float64 = 1
	for ; x > 0; x-- {
		y = y * x
	}
	return y
}
