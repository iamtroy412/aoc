package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func isEnginePartInSubLine(s string) bool {
    for _, c := range s {
        if c != '.' && !unicode.IsDigit(c) {
            return true
        }
    }
    
    return false
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

    var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, line)
    }

    var result int
    rg := regexp.MustCompile(`\d+`)

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if !unicode.IsDigit(rune(grid[i][j])) {
                continue
            }

            findNum := rg.FindString(grid[i][j:])
            num, _ := strconv.Atoi(findNum)
            // num = 467

            front := j
            if j > 0 {
                front = j - 1
            }

        }
    }


	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

}
