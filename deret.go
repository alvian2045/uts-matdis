package main

import "fmt"

func Deret1() {
	n := 12

	a := make([]int, n+1)

	// suku awal
	a[0] = 1
	a[1] = 1

	// hitung suku berikutnya
	for i := 2; i <= n; i++ {
		a[i] = 4*a[i-1] + 1*a[i-2]
	}

	// tampilkan semua suku dari a0 sampai an
	for i := 0; i <= n; i++ {
		fmt.Println("Suku ke", i, "=", a[i])

	}
	fmt.Printf("\nHASIL AKHIR Suku ke-%d: %d\n", n, a[n])

}
