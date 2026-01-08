package main

import (
	"fmt"
	"math"
)

func Deret2() {
	var a float64 = 10.00
	var r float64 = 0.2
	var n float64 = 8.0

	sn := a * (1 - math.Pow(r, n)) / (1 - r)
	s_takhingga := a / (1 - r)

	fmt.Println("Sum Berhingga S(8)\t:", sn)
	fmt.Println("Sum Tak Hingga S(inf)\t:", s_takhingga)

	persentase := float64(sn/s_takhingga) * (100 / 100)

	fmt.Printf("Persentase Kedekatan\t: %.f%%", persentase)

	fmt.Println("\nPersentase Kedekatan\t:", persentase, "%")

}
