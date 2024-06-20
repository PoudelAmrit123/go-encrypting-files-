package main

import (
	"bufio"
	"fmt"
	"os"
)

func encrypt(key int, plainText string) (result string) {
	return

}

func decpryt(key int, encryptedText string) (result string) {
	return
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	//Get the user input for the encryption
	fmt.Println("Enter the plain text for encryption")
	plainText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	encrypted := encrypt(5, plainText)
	fmt.Println("The plain text is ", encrypted)

	decryptedText := decpryt(5, encrypted)
	fmt.Println("The plain text is ", decryptedText)

}
