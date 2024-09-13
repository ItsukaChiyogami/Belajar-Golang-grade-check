package main

import "fmt"

func main() {

	var nilai int

	fmt.Println("Masukkan Nilai Murid")
	fmt.Scanln(&nilai)

	if nilai >= 90 {

		fmt.Println("Grade A")

	} else if nilai >= 80 {

		fmt.Println("Grade B")

	} else if nilai >= 70 {

		fmt.Println("Grade C")

	} else if nilai >= 60 {

		fmt.Println("Grade D")

	} else {

		fmt.Println("Grade F")

	}

}
