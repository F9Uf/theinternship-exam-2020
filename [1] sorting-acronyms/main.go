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

	sort.SliceStable(result, func(i, j int) bool {
		return len(result[i]) > len(result[j]) || (len(result[i]) == len(result[j])) && (result[i] < result[j])
	})

	for _, e := range result {
		fmt.Println(e)
	}

}

func getAcronyms(line string) string {
	var char []rune

	cutSpace := strings.Join(strings.Fields(strings.TrimSpace(line)), " ")
	words := strings.Split(cutSpace, " ")

	for _, word := range words {
		firstChar := []rune(word)[0]
		if unicode.IsUpper(firstChar) {
			char = append(char, firstChar)
		}
	}

	return string(char)
}
