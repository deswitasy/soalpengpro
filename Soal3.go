package main

import (
	"fmt"
)

func main() {
	// Meminta input dari pengguna
	var x int
	fmt.Print("Masukkan bilangan bulat positif x (<= 999): ")
	fmt.Scan(&x)

	// Memeriksa apakah x sesuai dengan batasan yang diberikan
	if x < 0 || x > 999 {
		fmt.Println("Input tidak valid. Masukkan bilangan bulat positif x (<= 999).")
		return
	}

	// Menghitung digit pertama, kedua, dan ketiga dari x
	d1 := x / 100       // Digit pertama (ratusan)
	d2 := (x / 10) % 10 // Digit kedua (puluhan)
	d3 := x % 10        // Digit ketiga (satuan)

	// Menampilkan hasil
	fmt.Printf("Digit pertama: %d\n", d1)
	fmt.Printf("Digit kedua  : %d\n", d2)
	fmt.Printf("Digit ketiga : %d\n", d3)
}
