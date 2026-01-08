package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Matrix1() {
	rand.Seed(time.Now().UnixNano())

	ordo := 2

	matrixA := make([][]int, ordo)
	matrixB := make([][]int, ordo)
	perkalian := make([][]int, ordo)

	for i := 0; i < ordo; i++ {
		matrixA[i] = make([]int, ordo)
		matrixB[i] = make([]int, ordo)
		perkalian[i] = make([]int, ordo)

		for j := 0; j < ordo; j++ {

			matrixA[i][j] = rand.Intn(20) + 1
			matrixB[i][j] = rand.Intn(20) + 1
		}
	}
	for i := 0; i < ordo; i++ {
		for j := 0; j < ordo; j++ {
			for k := 0; k < ordo; k++ {
				perkalian[i][j] += matrixA[i][k] * matrixB[k][j]

			}
		}
	}

	hasil_trace := perkalian[0][0] + perkalian[1][1]

	// Output
	fmt.Println("Matrix A:", matrixA)
	fmt.Println("Matrix B:", matrixB)
	fmt.Println("Hasil matriks (R):", perkalian)
	fmt.Println("Trace = ", hasil_trace)

}
