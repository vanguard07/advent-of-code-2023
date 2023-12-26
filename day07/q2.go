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

	jokerCount, jokerExists := charMap[rune('J')]

	if (jokerCount == 5) {
		return "FIVE"
	}

	oneCount, twoCount, threeCount, fourCount, fiveCount := 0, 0, 0, 0, 0
	for key, count := range charMap {
		if key == rune('J') {
			continue
		}
		if (count == 1) {
			oneCount += 1
		} else if (count == 2) {
			twoCount += 1
		} else if (count == 3) {
			threeCount += 1
		} else if (count == 4) {
			fourCount += 1
		} else {
			fiveCount += 1
		}
	}
	

	if (jokerExists) {
		if (fiveCount == 1 || fourCount == 1) {
			return "FIVE"
		} else if (threeCount == 1) {
			if (jokerCount == 2) {
				return "FIVE"
			} else {
				return "FOUR"
			}
		} else if (twoCount == 2) {
			return "FULL"
		} else if (twoCount == 1) {
			if (jokerCount == 3) {
				return "FIVE"
			} else if (jokerCount == 2) {
				return "FOUR"
			} else {
				return "THREE"
			}
		} else if (oneCount == 4) {
			return "ONEP"
		} else if (oneCount == 3) {
			return "THREE"
 		} else if (oneCount == 2) {
			return "FOUR"
		} else if (oneCount == 1) {
			return "FIVE"
		}
	} else {
		if (len(charMap) == 1) {
			return "FIVE"
		} else if (len(charMap) == 4) {
			return "ONEP"
		} else if (len(charMap) == 3) {
			if (threeCount == 1) {
				return "THREE"
			} else {
				return "TWOP"
			}
		} else if (len(charMap) == 2) {
			if (fourCount == 1 && oneCount == 1) {
				return "FOUR"
			} else {
				return "FULL"
			}
		} else {
			return "HIGH"
		}
	}
	return "FULL"
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
	cardMap[rune('J')] = 1
	cardMap[rune('2')] = 2
	cardMap[rune('3')] = 3
	cardMap[rune('4')] = 4
	cardMap[rune('5')] = 5
	cardMap[rune('6')] = 6
	cardMap[rune('7')] = 7
	cardMap[rune('8')] = 8
	cardMap[rune('9')] = 9
	cardMap[rune('T')] = 10
	cardMap[rune('Q')] = 11
	cardMap[rune('K')] = 12
	cardMap[rune('A')] = 13
	
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