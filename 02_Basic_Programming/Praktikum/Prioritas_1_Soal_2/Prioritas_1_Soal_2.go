package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func evenOdd(bilangan int) {

	if bilangan%2 == 0 {
		fmt.Printf("\nBilangan %d adalah genap\n", bilangan)
	} else {
		fmt.Printf("\nBilangan %d adalah ganjil\n", bilangan)
	}

	tryAgain()
}

func main()  {
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

	evenOdd(bilangan)
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