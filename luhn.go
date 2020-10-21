package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	checkCard := readerCard()
	// run algorithm
	if luhn(checkCard) {
		fmt.Printf("Card number %s exist !!! \n", checkCard)
		return
	}

	fmt.Printf("Card number %s doesn't exist !!! \n", checkCard)
}

// Luhn valid card using the Luhn Algorithm
// Return isvalid bool
func luhn(card string) bool {
	var (
		isSecond bool
		sum      int
	)

	// valid the lenght of string card
	if int(len(card)) < 14 || int(len(card)) > 19 {
		return false
	}

	// sum each digit on the string card
	for i := int(len(card) - 1); i >= 0; i-- {
		num, _ := strconv.Atoi(string(card[i]))
		digit := int(num)

		// multiplication the digit by 2 if this is the second digit in the string card
		if isSecond {
			digit = (digit * 2)
			if digit > 9 {
				digit -= 9
			}
		}

		// Take the sum of all digits
		sum += digit

		// change the flag second digit
		isSecond = !isSecond
	}

	// If the total modulo 10 is equal to 0 (if the total ends in zero)
	// then the number is valid according to the Luhn formula; otherwise it is not valid.
	return (sum%10 == 0)
}

// Read card number from the console
// Return card number string
func readerCard() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input number card")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)

	return text

}
