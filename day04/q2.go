package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	file, err := os.Open("day04/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	ans, lineNumber := 0, 0
	matchingMap := make(map[int]int)
	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		start, matches := 0, 0
		m := make(map[int]bool)
		for (string(line[start]) != "|") {
			if (line[start] >= 48 && line[start] <= 57) {
				num := 0
				for (line[start] >= 48 && line[start] <= 57) {
					num = num * 10 + (int(line[start]) - 48)
					start += 1
				}
				m[num] = true
			} else {
				start += 1;
			}
		}
		for start < len(line) {
			if (line[start] >= 48 && line[start] <= 57) {
				num := 0
				for (start < len(line) && line[start] >= 48 && line[start] <= 57) {
					num = num * 10 + (int(line[start]) - 48)
					start += 1
				}
				_, present := m[num]
				if (present) {
					matches += 1
				}
			} else {
				start += 1;
			}
		}
		matchingMap[lineNumber] = matches
	}
	file.Close()
	cardCounts := make([]int, lineNumber)
	for i := range cardCounts {
		cardCounts[i] = 1
	}
	for i := range cardCounts {
		_, present := matchingMap[i + 1]
		if (present && matchingMap[i + 1] > 0) {
			for j := i + 1; j < i + matchingMap[i + 1] + 1; j++ {
				cardCounts[j] += cardCounts[i]
			}
		}
	}
	for _, val := range cardCounts {
		ans += val
	}
	fmt.Println(ans)
}