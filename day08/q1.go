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

func main() {
	file, err := os.Open("day08/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var navigateSequence string
	if (scanner.Scan()) {
		navigateSequence = scanner.Text()
	}

	nodeMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			continue
		}

		keyValues := strings.Split(line, "=")
		key := strings.TrimSpace(keyValues[0])
		values := strings.Split(strings.Trim(strings.TrimSpace(keyValues[1]), "()"), ", ")
		for _, val := range values {
			nodeMap[key] = append(nodeMap[key], val)
		}
	}

	start, index, res := "AAA", 0, 0
	for start != "ZZZ" {
		if (navigateSequence[index] == 'R') {
			start = nodeMap[start][1]
		} else {
			start = nodeMap[start][0]
		}
		index = (index + 1) % len(navigateSequence)
		res += 1
	}
	
	file.Close()
	fmt.Println(res)
}