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

func strToList(str string) []string {
	str = strings.Trim(str, " \r\n")
	return strings.Split(str, " ")
}

func genDisplayList(numberOfProblem int) []string {
	var result []string
	for i := 0; i < numberOfProblem; i++ {
		result = append(result, "_")
	}
	return result
}

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
