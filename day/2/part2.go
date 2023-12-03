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
        gameSplit := strings.Split(line, ":")
        sets := getSets(gameSplit[1])
        minColors := map[string]int{
            "red": 0,
            "blue": 0,
            "green": 0,
        }
        for _, pull := range sets {
            colors := getColors(pull)
            if colors["red"] > minColors["red"] {
                minColors["red"] = colors["red"]
            }
            if colors["blue"] > minColors["blue"] {
                minColors["blue"] = colors["blue"]
            }
            if colors["green"] > minColors["green"] {
                minColors["green"] = colors["green"]
            }
        }
        power := minColors["red"] * minColors["blue"] * minColors["green"]
        sum += power
	}

    fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
