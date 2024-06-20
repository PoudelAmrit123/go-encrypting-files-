package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"

	// "os"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	key = key % len(runes)

	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOverLetters)
}

func encrypt(key int, plainText string) (result string) {

	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(hashLetter[letterPosition])
			return r

		}
		return r

	}

	strings.Map(findOne, plainText)
	return hashedString

}

func decpryt(key int, encryptedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(originalLetter[letterPosition])
			return r
		}
		return r

	}

	strings.Map(findOne, encryptedText)
	return hashedString
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	// Get the user input for the encryption
	fmt.Println("Enter the plain text for encryption")
	plainText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	plainText = strings.TrimSuffix(plainText, "\n")
	plainText = strings.ToUpper(plainText)
	// plainText := "HELLOWORLD"

	encryptedText := encrypt(5, plainText)
	fmt.Println("The encrypted text is ", encryptedText)

	decryptedText := decpryt(5, encryptedText)
	fmt.Println("The decrypted  text is ", decryptedText)

}
