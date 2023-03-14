package ceasarcipher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Path(input string) {
	// encode or decode
	var encodeChoice bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Type 'e' for encode and 'd' for decode > ")

	encodeChoiceInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	encodeChoiceInput = strings.TrimSuffix(encodeChoiceInput, "\n")
	if encodeChoiceInput == "d" {
		encodeChoice = false
	} else if encodeChoiceInput == "e" {
		encodeChoice = true
	} else {
		fmt.Println("Invalid choice")
		return
	}

	// key
	fmt.Print("Type your key (number) > ")
	keyReader := bufio.NewReader(os.Stdin)
	key, err := keyReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	key = strings.TrimSuffix(key, "\n")

	// message
	fmt.Print("Type your message > ")
	msgReader := bufio.NewReader(os.Stdin)
	msg, err := msgReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred")
		return
	}
	msg = strings.TrimSuffix(msg, "\n")

	CeasarCipher(msg, key, encodeChoice)
}
