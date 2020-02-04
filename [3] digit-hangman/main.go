package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	numberOfProblem := 12
	numberForGuess := 5

	fmt.Println("Enter The problem of hangman: ")
	problem, _ := reader.ReadString('\n')

	if list := strToList(problem); len(list) == numberOfProblem {
		displayList := genDisplayList(numberOfProblem)
		score := 0

		// loop for get the input number that user guess
		for i := 0; i < numberForGuess; i++ {
			printList(displayList)
			fmt.Printf("guess >> ")
			guessNumber, _ := reader.ReadString('\n')
			guessNumber = strings.Trim(guessNumber, " \r\n")

			score += getListToDisplay(list, &displayList, guessNumber)
		}
		printList(displayList)
		fmt.Println("Your score: ", score)
	} else {
		fmt.Println("Your problem size isn't equal 12!!")
	}

}

// strToList for split the string by space to the list
func strToList(str string) []string {
	str = strings.Trim(str, " \r\n")
	return strings.Split(str, " ")
}

// generate list for display that consist of "_"
func genDisplayList(numberOfProblem int) []string {
	var result []string
	for i := 0; i < numberOfProblem; i++ {
		result = append(result, "_")
	}
	return result
}

/*
	check the number that user guess is in the problem list?
	calculate score if the right guess
	but it's wil not calculate the duplicate answer
*/
func getListToDisplay(problemList []string, displayList *[]string, guessNumber string) int {
	score := 0
	isWrong := true
	for i := 0; i < len(problemList); i++ {
		if guessNumber == problemList[i] {
			if (*displayList)[i] == "_" {
				score++
			}

			(*displayList)[i] = guessNumber
			isWrong = false
		}
	}
	if isWrong {
		(*displayList) = append(*displayList, guessNumber)
	}
	return score
}

func printList(list []string) {
	fmt.Println(strings.Join(list, " "))
}
