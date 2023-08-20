package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func faktorBilangan(bilangan int) {

	if bilangan > 0 {
		for i := 1; i <= bilangan; i++ { //rumus faktorial
			if bilangan%i == 0 {
				fmt.Println(i)
			}
		}
	}else{
		main()
	}

	tryAgain()
}

func main(){
	var input string
	var bilangan int

	clear()
	fmt.Print("Soal 2 (Prioritas)\n\n")

	for {
		fmt.Print("Masukkan nilai : ")
		_, err := fmt.Scanln(&input)

		bilangan, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka.")
			continue
		}
		break
	}

	faktorBilangan(bilangan)
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