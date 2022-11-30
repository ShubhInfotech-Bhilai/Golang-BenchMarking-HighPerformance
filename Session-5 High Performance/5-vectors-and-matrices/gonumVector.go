package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func matrixPrint(m mat.Matrix) {
	formattedMatrix := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", formattedMatrix)
}

func main() {
	v := mat.NewVecDense(4, []float64{0, 1, 2, 3})
 a1:=[]float64{0, 1, 2, 3}
 a2:=[]float64{10, 11, 12, 13}
fmt.Println(stat.Mean(a1,a2) )

matrixPrint(v)


}
