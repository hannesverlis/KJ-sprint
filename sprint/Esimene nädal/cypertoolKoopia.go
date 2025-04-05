package main

import (
	"fmt"
	"strings"
)

func main() {

	// basic user greeting
	fmt.Println("Welcome to the Cypher Tool!") //greet the user

	toEncrypt, encoding, message := getInput() //run the getInput function to get toEncrypt, encoding, message
	result := ""
	toEncryptStr := ""
	cypherStr := ""

	if toEncrypt == true {
		toEncryptStr = "Encrepted"
		switch encoding {
		case "1":
			cypherStr = "rot13"
			result = encrypt_rot13(message) //choose rot13
		case "2":
			cypherStr = "reverse"
			result = encrypt_reverse(message) //choose reverse
		case "3":
			cypherStr = "third chypher"
			result = encrypt_thirdCypher(message) //choose third Cypher
		}
	} else { //when encrypt is falce , user selects descrypt
		toEncryptStr = "Descrypted"
		switch encoding {
		case "1":
			cypherStr = "rot13"
			result = decrypt_rot13(message) //choose rot13
		case "2":
			cypherStr = "reverse"
			result = decrypt_reverse(message) //choose reverse
		case "3":
			cypherStr = "third chypher"
			result = decrypt_thirdCypher(message) //choose third Cypher
		}
	}
	fmt.Printf("\v%v message using %v: %v\n", toEncryptStr, cypherStr, result)
}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {

	var scanner string //new variable for conventing user input "1" or "2" to boolean value
	//ask for selection

	fmt.Println("\vSelect operation (1/2):\n") //added \v and \n for formating it likt was showsd in the task
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	fmt.Scanln(&scanner) // converting scanner value to boolean value
	for !toEncrypt {     //if we dont reach toEncrypt, loop is going
		if scanner == "1" {
			toEncrypt = true
			break //loop ends
		} else if scanner == "2" {
			toEncrypt = false
			break //loop ends
		} else {
			fmt.Println("\vYou can select only 1/2. Please try again\n")
			fmt.Scan(&scanner)

		}
	}
	fmt.Println("\vSelect cypher (1-3):\n") //added \v and \n for formating it likt was showsd in the task
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. Third cypher")
	fmt.Scanln(&encoding)

	fmt.Printf("\vEnter the message:\n")
	fmt.Scanln(&message)
	message = strings.TrimSpace(message) //removes whitespaces
	return toEncrypt, encoding, message
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {

	result := make([]rune, len(s)) //a slice result of type rune is created with a length equal to the length of string s.
	for i, c := range s {          //The for loop loops through each character of the string s
		if isLetter(c) { //This condition checks whether the current character is a letter (of the Latin alphabet).
			if isLower(c) { //If the character is a letter, check to see if it is lowercase (small)
				result[i] = 'z' - (c - 'a')
			} else {
				result[i] = 'Z' - (c - 'A')
			}
		} else {
			result[i] = c
		}
	}
	return string(result)

	return s
}

func rotate(c rune, start rune, end rune) rune {
	return start + (c-start+13)%(end-start+1)
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	result := make([]rune, len(s))
	for i, c := range s {
		if isLetter(c) {
			if isLower(c) {
				result[i] = rotate(c, 'a', 'z')
			} else {
				result[i] = rotate(c, 'A', 'Z')
			}
		} else {
			result[i] = c
		}
	}
	return string(result)
	return s
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
	result := make([]rune, len(s))
	for i, c := range s {
		if isLetter(c) {
			if isLower(c) {
				result[i] = rotate(c, 'a', 'z')
			} else {
				result[i] = rotate(c, 'A', 'Z')
			}
		} else {
			result[i] = c
		}
	}
	return string(result)
	return s
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	result := make([]rune, len(s)) //a slice result of type rune is created with a length equal to the length of string s.
	for i, c := range s {          //The for loop loops through each character of the string s
		if isLetter(c) { //This condition checks whether the current character is a letter (of the Latin alphabet).
			if isLower(c) { //If the character is a letter, check to see if it is lowercase (small)
				result[i] = 'z' - (c - 'a')
			} else {
				result[i] = 'Z' - (c - 'A')
			}
		} else {
			result[i] = c
		}
	}
	return string(result)

	return s
}

// Encrypt the message with
func encrypt_thirdCypher(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)

	return s
}

// Decrypt the message with
func decrypt_thirdCypher(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)

	return s
}
