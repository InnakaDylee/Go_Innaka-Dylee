package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"bufio"
	"strings"
)

func palindrome(nama string) {
	var verif bool

	nameLower := strings.ToLower(nama)

	for i := 0; i < len(nama)/2; i++ { //verif palindrome
		if nameLower[i] == nameLower[len(nama)-i-1] {
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

	tryAgain()
}

func main(){
	var nama string

	reader := bufio.NewReader(os.Stdin)
	clear()
	fmt.Print("Soal 1 (Explorasi)\n\n")

	fmt.Print("Masukkan nama : ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	palindrome(nama)
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