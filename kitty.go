package main

import (
	"bufio"
	"fmt"
	kt "kitty/files"
	"os"
	"os/exec"
	"runtime"
)

var logo string = `
█  █▀ ▄█    ▄▄▄▄▀ ▄▄▄▄▀ ▀▄    ▄ 
█▄█   ██ ▀▀▀ █ ▀▀▀ █      █  █  
█▀▄   ██     █     █       ▀█   
█  █  ▐█    █     █        █    
  █    ▐   ▀     ▀       ▄▀     
 ▀ 
`

// clear console function
func Clear() {
	system := runtime.GOOS
	if system == "windows" {
		console := exec.Command("cmd", "/c", "cls")
		console.Stdout = os.Stdout
		console.Run()
	} else {
		console := exec.Command("clear")
		console.Stdout = os.Stdout
		console.Run()
	}
}

// run project
func main() {
	for {
		Clear()
		// Home
		fmt.Printf(`
	%v
		<§> Enter your text for Encrypt

		<1> Encrypt

		<2> Decrypt

		<3> About
		
		<4> Exit

		┌─[Kitty~ #Home]
		└╼ <†> `, logo)

		// First Input (Home page)
		var FirstPut int
		fmt.Scanln(&FirstPut)

		// Encrypt
		if FirstPut == 1 {
			Clear()
			fmt.Printf(`
	%v
		<§> Enter your text for Encrypt

		┌─[Kitty~ #Home/Encrypt]
		└╼ § `, logo)
			// Input String for encrypt
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			PutEncrypt := scanner.Text()

			// Encrypt cal
			encrypt, err := kt.Encrypt(PutEncrypt, kt.MySecret)
			if err != nil {
				fmt.Println("error encrypting your classified text: ", err)
			}
			Clear()
			res := "your text has EnCrypted!\nResualt: " + encrypt
			fmt.Printf(`
		%v
%v

		<§> Enter your choice:

		<1> Back To Home

		<2> Exit

		┌─[Kitty~ #Home/Encrypt]
		└╼ § `, logo, res)

			// continue or exit from console
			var SecoundPut int
			fmt.Scanln(&SecoundPut)

			if SecoundPut == 1 {
				continue
			} else {
				fmt.Println("Bye!")
				os.Exit(0)
			}

			// Decrypt
		} else if FirstPut == 2 {
			Clear()
			fmt.Printf(`
	%v
		<§> Enter your text for Decrypt

		┌─[Kitty~ #Home/Decrypt]
		└╼ § `, logo)

			// Input Encrypt text for Decrypt
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			PutDecrypt := scanner.Text()

			// Decrypt cal
			decrypt, err := kt.Decrypt(PutDecrypt, kt.MySecret)
			if err != nil {
				fmt.Println("error: ", err)
			}
			Clear()
			res := "your text has DeCrypted!\nResualt: " + decrypt
			fmt.Printf(`
		%v
%v

		<§> Enter your choice:

		<1> Back To Home

		<2> Exit

		┌─[Kitty~ #Home/Decrypt]
		└╼ § `, logo, res)

			// continue or exit from console
			var SecoundPut int
			fmt.Scanln(&SecoundPut)
			if SecoundPut == 1 {
				continue
			} else {
				fmt.Println("Bye!")
				os.Exit(0)
			}

			// About
		} else if FirstPut == 3 {
			Clear()
			fmt.Printf(`
%v
<§> Kitty can encrypt or even decrypt your texts in a very cute way.
<!> Kitty love you :D

<$> Project Address: https://github.com/TheFr3Y/Kitty/
			`, logo)
			fmt.Printf(`

		<§> Enter your choice:

		<1> Back To Home

		<2> Exit

		┌─[Kitty~ #Home/About]
		└╼ § `)
			var SecoundPut int
			fmt.Scanln(&SecoundPut)
			if SecoundPut == 1 {
				continue
			} else {
				fmt.Println("Bye!")
				os.Exit(0)
			}

			// Exit
		} else {
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}

}
