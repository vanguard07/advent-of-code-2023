package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getType (str string) string {
	charMap := make(map[rune]int)

	for _, char := range str {
		charMap[char]++
	}

	if (len(charMap) == 5) {
		return "HIGH"
	} else if (len(charMap) == 4) {
		return "ONEP"
	} else if (len(charMap) == 3) {
		oneCount, twoCount, threeCount := 0, 0, 0
		for _, count := range charMap {
			if (count == 1) {
				oneCount += 1
			} else if (count == 2) {
				twoCount += 1
			} else if (count == 3) {
				threeCount += 1
			}
		}
		if (threeCount == 1) {
			return "THREE"
		} else {
			return "TWOP"
		}
	} else if (len(charMap) == 2) {
		oneCount, twoCount, threeCount, fourCount := 0, 0, 0, 0
		for _, count := range charMap {
			if (count == 1) {
				oneCount += 1
			} else if (count == 2) {
				twoCount += 1
			} else if (count == 3) {
				threeCount += 1
			} else {
				fourCount += 1
			}
		}
		if (fourCount == 1 && oneCount == 1) {
			return "FOUR"
		} else {
			return "FULL"
		}
	}
	return "FIVE"
}

func main() {
	file, err := os.Open("day07/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	handTypeMap := make(map[string][]string)
	handCostMap := make(map[string]int)
	rankMap := make(map[int]string)
	rankMap[0] = "HIGH"
	rankMap[1] = "ONEP"
	rankMap[2] = "TWOP"
	rankMap[3] = "THREE"
	rankMap[4] = "FULL"
	rankMap[5] = "FOUR"
	rankMap[6] = "FIVE"

	cardMap := make(map[rune]int)
	cardMap[rune('2')] = 1
	cardMap[rune('3')] = 2
	cardMap[rune('4')] = 3
	cardMap[rune('5')] = 4
	cardMap[rune('6')] = 5
	cardMap[rune('7')] = 6
	cardMap[rune('8')] = 7
	cardMap[rune('9')] = 8
	cardMap[rune('T')] = 9
	cardMap[rune('J')] = 10
	cardMap[rune('Q')] = 11
	cardMap[rune('K')] = 12
	cardMap[rune('A')] = 13

	// fmt.Println(cardMap)
	
	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		
		hand := entry[0]
		cost, _ := strconv.Atoi(entry[1])
		handCostMap[hand] = cost

		handType := getType(hand)
		handTypeMap[handType] = append(handTypeMap[handType], hand)
	}

	for key, values := range handTypeMap {
		sort.Slice(values, func(i, j int) bool {
			for idx := 0; idx < min(len(values[i]), len(values[j])); idx++ {
				if (values[i][idx] != values[j][idx]) {
					fmt.Println(values[i], cardMap[rune(values[i][idx])], values[j], cardMap[rune(values[j][idx])])
					return cardMap[rune(values[i][idx])] < cardMap[rune(values[j][idx])]
				}
			}
			return true
		})
		handTypeMap[key] = values
	}

	rank, ans := 1, 0
	for i := 0; i < len(rankMap); i++ {
		for _, value := range handTypeMap[rankMap[i]] {
			fmt.Println(rank, value, handCostMap[value])
			ans += rank * handCostMap[value]
			rank++
		} 
	}

	file.Close()
	fmt.Println(ans)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}