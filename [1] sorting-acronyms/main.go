package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var result []string

	fmt.Printf("Enter number of string: ")

	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = strings.Trim(num, "\r\n")
	n, _ := strconv.Atoi(num)

	for i := 0; i < n; i++ {

		fmt.Printf(">> ")
		line, _ := reader.ReadString('\n')

		acronym := getAcronyms(line)
		result = append(result, acronym)
	}

	// sort by length of acronyms or if equal it will sort by Alphabetical order
	sort.SliceStable(result, func(i, j int) bool {
		return len(result[i]) > len(result[j]) || (len(result[i]) == len(result[j])) && (result[i] < result[j])
	})

	for _, e := range result {
		fmt.Println(e)
	}

}

/*
	getAcronyms
		require line of string, full name
		and return acronym name
*/
func getAcronyms(line string) string {
	var char []rune

	// cut dobule space between word
	cutSpace := strings.Join(strings.Fields(strings.TrimSpace(line)), " ")
	// split the string into the list of word
	words := strings.Split(cutSpace, " ")

	// loop for check uppercase
	for _, word := range words {
		firstChar := []rune(word)[0]
		if unicode.IsUpper(firstChar) {
			char = append(char, firstChar)
		}
	}

	return string(char)
}
