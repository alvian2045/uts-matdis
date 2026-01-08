package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Matrix2() {
	rand.Seed(time.Now().UnixNano())

	ordo := 3
	matrix := make([][]int, ordo)

	for i := range matrix {
		matrix[i] = make([]int, ordo)
		for j := 0; j < ordo; j++ {
			matrix[i][j] = rand.Intn(20) + 1

			for i := 0; i < ordo; i++ {
				for j := 0; j < ordo; j++ {

				}
			}
		}
	}

	fmt.Println("matrix M (awal) : ", matrix)

	matrix[0], matrix[ordo-1] = matrix[ordo-1], matrix[0]
	fmt.Println("menukar baris 0 dan ", ordo-1)
	fmt.Println("Matrix M terkini : ", matrix)

	max := matrix[0][0]
	baris, kolom := 0, 0

	for i := 0; i < ordo; i++ {
		for j := 0; j < ordo; j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
				baris = i
				kolom = j
			}
		}
	}

	fmt.Printf(
		"Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n",
		max, baris, kolom,
	)
}
