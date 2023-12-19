package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	file, err := os.Open("day02/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	num, ans := 0, 0
	for scanner.Scan() {
		num += 1
		line := scanner.Text()
		gameRecord := strings.Split(strings.Split(line, ":")[1], ";")
		isValidRecord := true
		for _, records := range gameRecord {
			record := strings.Split(strings.TrimSpace(records), ", ")
			for _, draw := range record {
				numberAndColor := strings.Split(strings.TrimSpace(draw), " ")
				x, err := strconv.ParseInt(numberAndColor[0], 10, 64)
				if err != nil {
					panic(err)
				}
				if ((numberAndColor[1] == "red" &&  x > 12) || 
				(numberAndColor[1] == "green" &&  x > 13) || 
				(numberAndColor[1] == "blue" &&  x > 14)) {
					isValidRecord = false
				}
			}
			if (!isValidRecord) {
				break
			}
		}
		if (isValidRecord) {
			ans += num
		}
	}
	file.Close()
	fmt.Println(ans)
}