package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func segitigaAsterik(row int) {

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
		main()
	}

	tryAgain()
}

func main(){
	var input string
	var row int

	clear()
	fmt.Print("Soal 1 (Prioritas 2)\n\n")

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

	segitigaAsterik(row)
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