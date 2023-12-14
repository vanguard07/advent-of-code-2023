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
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameRecord := strings.Split(strings.Split(line, ":")[1], ";")
		red, green, blue := 0, 0, 0
		for _, records := range gameRecord {
			record := strings.Split(strings.TrimSpace(records), ", ")
			for _, draw := range record {
				numberAndColor := strings.Split(strings.TrimSpace(draw), " ")
				x, err := strconv.ParseInt(numberAndColor[0], 10, 64)
				if err != nil {
					panic(err)
				}
				if (numberAndColor[1] == "red") {
					red = max(red, int(x))
				}
				if (numberAndColor[1] == "green") {
					green = max(green, int(x))
				}
				if (numberAndColor[1] == "blue") {
					blue = max(blue, int(x))
				}
			}
		}
		ans += red * green * blue
	}
	file.Close()
	fmt.Println(ans)
}