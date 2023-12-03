package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first, last int
		line := scanner.Text()
		for _, c := range line {
			if unicode.IsDigit(c) {
				first = int(c)
				break
			}
		}

		r := []rune(line)
		for i := len(r) - 1; i >= 0; i-- {
			if unicode.IsDigit(r[i]) {
				last = int(r[i])
				break
			}
		}
		t, _ := strconv.Atoi(string(first) + string(last))

		sum += t
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
