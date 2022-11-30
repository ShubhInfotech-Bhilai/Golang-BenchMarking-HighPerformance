package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{5, 3, 10, 1, 6, 4, 8, 7, 2})
	matrixPrint(a)
	
	fmt.Println(a.At(0,2)) //( row ,col)
	/*
	matrixPrint(a.T())
	fmt.Println(a.Dims())
	*/
	

}

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}
