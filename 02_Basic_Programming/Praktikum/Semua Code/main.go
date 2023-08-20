package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"bufio"
	"strings"
)

// No. 1 Prioritas 1
func trapesium() {
	var inputAtas, inputAlas, inputTinggi string
	var atas, alas, tinggi float64

	clear()
	fmt.Print("Soal 1\n\n")

	for {
		fmt.Print("Masukkan nilai atas : ")
		_, err := fmt.Scanln(&inputAtas)

		atas, err = strconv.ParseFloat(inputAtas, 64)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan nilai alas : ")
		_, err := fmt.Scanln(&inputAlas)

		alas, err = strconv.ParseFloat(inputAlas, 64)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan nilai tinggi : ")
		_, err := fmt.Scanln(&inputTinggi)

		tinggi, err = strconv.ParseFloat(inputTinggi, 64)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	luas := (atas + alas) / 2 * tinggi //rumus luas trapesium1

	fmt.Printf("\nLuas Trapesium     : %f\n", luas)

	backHome()
}

// No. 2 Prioritas 1
func evenOdd() {
	var input string
	var bilangan int

	clear()
	fmt.Print("Soal 2\n\n")

	for {
		fmt.Print("Masukkan nilai bilangan : ")
		_, err := fmt.Scanln(&input)

		bilangan, err = strconv.Atoi(input)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	if bilangan%2 == 0 {
		fmt.Printf("\nBilangan %d adalah genap", bilangan)
	} else {
		fmt.Printf("\nBilangan %d adalah ganjil", bilangan)
	}

	backHome()
}

// No. 3 Prioritas 1
func gradeNilai() {
	var input string
	var nilai int

	fmt.Print("Soal 3\n\n")

	for {
		fmt.Print("Masukkan nilai : ")
		_, err := fmt.Scanln(&input)

		nilai, err = strconv.Atoi(input)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	if nilai >= 0 && nilai <= 100 { //cek apakah nilai kurang dari 0 dan lebih dari 100
		if nilai >= 80 && nilai <= 100 {
			fmt.Print("Nilai A")
		} else if nilai >= 65 && nilai <= 79 {
			fmt.Print("Nilai B")
		} else if nilai >= 50 && nilai <= 64 {
			fmt.Print("Nilai C")
		} else if nilai >= 35 && nilai <= 49 {
			fmt.Print("Nilai D")
		} else {
			fmt.Print("Nilai E")
		}
	} else {
		fmt.Println("Nilai Invalid")
		gradeNilai()
	}

	backHome()
}

// No. 4 Prioritas 1
func kelipatan() {

	clear()
	fmt.Print("Soal 4\n\n")

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 { //Kelipatan 3 & 5
			fmt.Print("fizzbuzz ")
		} else if i%3 == 0 { //Kelipatan 3
			fmt.Print("fizz ")
		} else if i%5 == 0 { //Kelipatan5
			fmt.Print("buzz ")
		} else {
			fmt.Printf("%d ", i)
		}

		//note: angka 15 30 45 60 75 90 yang habis dibagi 3 dan 5
		//tidak diperintah untuk menampilkan hal khusus sehingga
		//saya buat gabungan dari 3 dan 5 yaitu fibuzz
	}

	backHome()
}

// No. 1 Prioritas 2
func segitigaAsterik() {
	var input string
	var row int

	clear()
	fmt.Print("Soal 1 (Prioritas)\n\n")

	for {
		fmt.Print("Masukkan baris : ")
		_, err := fmt.Scanln(&input)

		row, err = strconv.Atoi(input)//error handling
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	if row > 0 {// validasi nilai baris tidak boleh 0
		for i := 0; i <= row; i++ {
			for j := 1; j <= row-i; j++ { //print kosong untuk space kosong dari kiri
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ { //print bintang diujung space kosong
				fmt.Print(" *")
			}
			fmt.Println(" ") //baris
		}
	}else{
		faktorBilangan()
	}


	backHome()
}

// No. 2 Prioritas 2
func faktorBilangan() {
	var input string
	var a int

	clear()
	fmt.Print("Soal 2 (Prioritas)\n\n")

	for {
		fmt.Print("Masukkan nilai : ")
		_, err := fmt.Scanln(&input)

		a, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	if a > 0 {
		for i := 1; i <= a; i++ { //rumus faktorial
			if a%i == 0 {
				fmt.Println(i)
			}
		}
	}else{
		faktorBilangan()
	}

	backHome()
}

// No. 1 Explorasi
func palindrome() {
	var nama string
	var verif bool

	reader := bufio.NewReader(os.Stdin)
	clear()
	fmt.Print("Soal 1 (Explorasi)\n\n")

	fmt.Print("Masukkan nama : ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	for i := 0; i < len(nama)/2; i++ { //verif palindrome
		if nama[i] == nama[len(nama)-i-1] {
			verif = true
		} else {
			verif = false
		}
	}
	fmt.Printf("Captured : %v\n", nama) //print input
	fmt.Print("Invert   : ")            //print invert

	for i := len(nama) - 1; i >= 0; i-- {
		fmt.Print(string(nama[i]))
	}

	if verif == true {
		fmt.Println("\nPalindrome")
	} else {
		fmt.Println("\nBukan Palindrome")
	}

	backHome()
}

func main() {
	var input int

	clear()
	fmt.Print("Halaman Awal\n\n")
	fmt.Println("1. Soal Prioritas 1")
	fmt.Println("2. Soal Prioritas 2")
	fmt.Println("3. Soal Eksplorasi")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&input)

	if input == 1 {
		var input string
		var a int

		clear()
		fmt.Print("Soal Prioritas 1\n\n")
		fmt.Println("1. Luas Trapesium")
		fmt.Println("2. Ganjil & Genap")
		fmt.Println("3. Grade Nilai")
		fmt.Println("4. Kelipatan 3 & Kelipatan 5")
		for {
			fmt.Print("Masukkan Pilihan : ")
			_, err := fmt.Scanln(&input)
	
			a, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan hanya angka.")
				continue
			}
			break
		}

		if a == 1 {
		} else if a == 2 {
			evenOdd()
		} else if a == 3 {
			clear()
			gradeNilai()
		} else if a == 4 {
			kelipatan()
		}

	} else if input == 2 {
		var input string
		var a int

		clear()
		fmt.Print("Soal Prioritas 2\n\n")
		fmt.Println("1. Segitiga Asterik")
		fmt.Println("2. Faktor Bilangan")
		for {
			fmt.Print("Masukkan Pilihan : ")
			_, err := fmt.Scanln(&input)
	
			a, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input tidak valid, masukkan hanya angka.")
				continue
			}
			break
		}

		if a == 1 {
			segitigaAsterik()
		} else if a == 2 {
			faktorBilangan()
		}
	} else if input == 3 {
		palindrome()
	} else {
		main()
	}
}

func backHome() {
	var back string

	fmt.Print("\nBalik ke Halaman Utama? (Y/N)")
	fmt.Scanln(&back)
	if back == "y" || back == "Y" { //home
		main()
	} else if back == "n" || back == "N" { //close
		fmt.Println("Terima Kasih!! ^_^")
	} else {
		backHome() //refresh invalid button
	}
}

func clear() {
	if runtime.GOOS == "windows" { // clear windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else { // clear linux, macOS, etc.
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
