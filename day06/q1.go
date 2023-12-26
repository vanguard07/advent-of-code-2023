package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day06/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	ans := 1
	var time, distance string
	if scanner.Scan() {
		time = scanner.Text()
	}
	if scanner.Scan() {
		distance = scanner.Text()
	}

	var times, distances []int
	
	numStrings := strings.Fields(strings.Split(time, ":")[1])
	for _, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		times = append(times, num)
	}
	
	numStrings = strings.Fields(strings.Split(distance, ":")[1])
	for _, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		distances = append(distances, num)
	}
	for i, t := range(times) {
		prev := 0
		count := 0
		for j := 1; j < t + 1; j++ {
			if (prev >= j * (t - j)) {
				break
			}
			if (distances[i] < j * (t - j)) {
				count += 1
			}
			prev = j * (t - j)
		}
		if (count > 0) {
			count *= 2
			if (t % 2 == 0) {
				count -= 1
			}
			ans *= count
		}
	}

	file.Close()
	fmt.Println(ans)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
