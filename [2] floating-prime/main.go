package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter Floating Prime: ")
	// read line until user type "0.0"
	for x, _ := bufio.NewReader(os.Stdin).ReadString('\n'); strings.Trim(x, "\r\n") != "0.0"; x, _ = bufio.NewReader(os.Stdin).ReadString('\n') {

		// trim string and remove "." out
		number := strings.Trim(x, "\r\n")
		number = strings.Join(strings.Split(number, "."), "")

		var result string
		result = "FALSE"

		// loop check prime of number until found the prime number
		// Loop sinc one floating point to 3 floating point
		for i := 2; i <= 4; i++ {
			substr := number[0:i]
			if x, err := strconv.Atoi(substr); err == nil {
				if isPrime(x) {
					result = "TRUE"
					break
				}
			}
		}

		fmt.Println(">> ", result)
		fmt.Printf("Enter Floating Prime: ")

	}
}

/*
	isPrime is a function to check prime of input number
*/
func isPrime(number int) bool {
	for i := 2; i <= int(math.Floor(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
