package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Himpunan2() {
	rand.Seed(time.Now().UnixNano())

	M := 8
	K := 18

	// Membuat himpunan S secara random
	S := make([]int, M)
	for i := 0; i < M; i++ {
		S[i] = rand.Intn(20) + 1 // angka random 1–20
	}

	fmt.Println("Set S:", S, "| Target K:", K)
	fmt.Println("\nSubset 2-Elemen (Sum > 18):")

	count := 0

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			sum := S[i] + S[j]
			if sum > K {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n",
					count, S[i], S[j], sum)
			}
		}
	}

	fmt.Println("\nTotal:", count, "Pasangan")
}
