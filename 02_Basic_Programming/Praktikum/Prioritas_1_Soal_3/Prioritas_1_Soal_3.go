package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func gradeNilai(nilai int) {

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
		main()
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

	gradeNilai(nilai)
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