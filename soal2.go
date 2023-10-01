package main

import (
	"fmt"
)

func calculateEquation(x float64) float64 {
	//  fungsi calculateEquation itu untuk yang menerima sebuah parameter berupa bilangan berkoma x
	numerator := x*x*x + 3*x
	denominator := x*x*x*x - 3*x*x + 4
	result := numerator / denominator
	return result
}

func main() {
	// mendeklarasikan variabel x untuk menyimpan input dari pengguna dengan fmt.Scan(). lalu,  memanggil fungsi calculateEquation dengan nilai x yang diinputkan oleh pengguna. Hasil perhitungan disimpan dalam variabel output.
	var x float64
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	output := calculateEquation(x)
	fmt.Println("Nilai f(x) adalah:", output)
}
