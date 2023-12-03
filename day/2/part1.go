package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSets(s string) []string {
    return strings.Split(s, ";")
}

func getColors(s string) map[string]int {
    cmap := map[string]int{}
    colors := strings.Split(s, ",")
    for _, c := range colors {
        c = strings.TrimSpace(c)
        ind := strings.Split(c, " ")
        amt, _ := strconv.Atoi(ind[0])
        cmap[ind[1]] += amt
    }
    return cmap
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

    sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        possible := true
        gameSplit := strings.Split(line, ":")
        id := strings.Split(gameSplit[0], " ")[1]
        sets := getSets(gameSplit[1])
        rTotal := 12
        bTotal := 14
        gTotal := 13
        for _, pull := range sets {
            colors := getColors(pull)
            if colors["red"] > rTotal || colors["blue"] > bTotal || colors["green"] > gTotal {
                possible = false
                break
            }

        }
        if possible {
            i, _ := strconv.Atoi(id)
            sum += i
        }
	}

    fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
