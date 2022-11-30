package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}
func main() {
	a := mat.NewDense(1, 3, []float64{0.10, 0.42, 0.37})
	b := mat.NewDense(3, 2, []float64{5, 8, 10, 6, 2, 3})
	matrixPrint(a)
	matrixPrint(b)

	var c mat.Dense
	c.Mul(a, b)
	result := mat.Formatted(&c, mat.Prefix(""), mat.Squeeze())
	fmt.Println(result)
}
