package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func trapesium(atas, alas, tinggi float64) {

	luas := (atas + alas) / 2 * tinggi //rumus luas trapesium1

	fmt.Printf("\nLuas Trapesium     : %f\n", luas)

	tryAgain()
}

func main(){
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
	
	trapesium(atas, alas, tinggi)
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