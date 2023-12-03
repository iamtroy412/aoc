package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	regex := regexp.MustCompile("(\\d|one|two|three|four|five|six|seven|eight|nine)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
        num := match(line, regex)
        sum += num
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}

func match(s string, r *regexp.Regexp) int {
	var matches []string
	for i := 0; i < len(s); i++ {
		m := r.FindAllString(s[i:], -1)
		matches = append(matches, m...)
	}

	if len(matches) > 0 {
        fnum := toInt(matches[0])
        lnum := toInt(matches[len(matches)-1])
        return fnum*10 + lnum
	}

	return 0
}

func toInt(s string) int {
	s2n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	if s2n[s] != 0 || s == "0" {
		return s2n[s]
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}
