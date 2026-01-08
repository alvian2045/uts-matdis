package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ada(data []int, x int) bool {
	for _, v := range data {
		if v == x {
			return true
		}
	}
	return false
}

func randomSet(jumlah, n int) []int {
	var hasil []int
	for len(hasil) < jumlah {
		x := rand.Intn(n) + 1
		if !ada(hasil, x) {
			hasil = append(hasil, x)
		}
	}
	return hasil
}

func Himpunan1() {
	rand.Seed(time.Now().UnixNano())

	n := 120
	A := randomSet(5, n)
	B := randomSet(5, n)
	C := randomSet(5, n)

	var A_xor_B []int
	var R []int
	var RGenap []int

	// A ⊕ B
	for _, a := range A {
		if !ada(B, a) {
			A_xor_B = append(A_xor_B, a)
		}
	}
	for _, b := range B {
		if !ada(A, b) {
			A_xor_B = append(A_xor_B, b)
		}
	}

	// (A ⊕ B) \ C
	for _, x := range A_xor_B {
		if !ada(C, x) {
			R = append(R, x)
		}
	}

	// A ∩ C
	for _, a := range A {
		if ada(C, a) {
			R = append(R, a)
		}
	}
	for _, x := range R {
		if x%2 == 0 {
			RGenap = append(RGenap, x)
		}
	}

	fmt.Println("n =", n)
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	fmt.Println("C =", C)
	fmt.Println("R =", R)
	fmt.Println("Hasil filter (Bil.Genap) =", RGenap)
}
