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
	for x, _ := bufio.NewReader(os.Stdin).ReadString('\n'); strings.Trim(x, "\r\n") != "0.0"; x, _ = bufio.NewReader(os.Stdin).ReadString('\n') {
		number := strings.Trim(x, "\r\n")
		number = strings.Join(strings.Split(number, "."), "")

		var result string
		result = "FALSE"

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

func isPrime(number int) bool {
	for i := 2; i <= int(math.Floor(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
