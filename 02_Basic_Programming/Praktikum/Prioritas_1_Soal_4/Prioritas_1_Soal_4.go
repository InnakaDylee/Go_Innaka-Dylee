package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func kelipatan(nilai int) {

	clear()
	fmt.Print("Soal 4\n\n")

	for i := 1; i <= nilai; i++ {

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

	tryAgain()
}

func main(){
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

	kelipatan(nilai)
}

func tryAgain() {
	var back string

	fmt.Print("\nCoba Lagi? (Y/N)")
	fmt.Scanln(&back)
	if back == "y" || back == "Y" { //home
		main()
	} else if back == "n" || back == "N" { //close
		fmt.Println("Terima Kasih!! ^_^")
	} else {
		tryAgain() //refresh invalid button
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