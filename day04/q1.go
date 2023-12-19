package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
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
	ans := 0
	for scanner.Scan() {
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
		if (matches > 0) {
			ans += int(math.Pow(2, float64(matches - 1)))
		}
	}
	file.Close()
	fmt.Println(ans)
}