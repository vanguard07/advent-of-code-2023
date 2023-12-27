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

func gcd(a, b int) int {
	if (b == 0) {
		return a
	}
	return gcd(b, a % b)
}

func lcm(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		n1, n2 := result, nums[i]
		gcd_val := gcd(n1, n2)
		result = (result * nums[i]) / gcd_val
	}

	return result
}


func checkIfAllEndWithZ (arr [] string) bool {
	count := 0
	for _, val := range arr {
		if (val[2] == 'Z') {
			count += 1
		}
	}
	
	fmt.Println(count, len(arr))
	return count != len(arr)
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
	var startArray []string
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			continue
		}

		keyValues := strings.Split(line, "=")
		key := strings.TrimSpace(keyValues[0])
		if (key[2] == 'A') {
			startArray = append(startArray, key)
		}
		values := strings.Split(strings.Trim(strings.TrimSpace(keyValues[1]), "()"), ", ")
		for _, val := range values {
			nodeMap[key] = append(nodeMap[key], val)
		}
	}

	var timeToZ []int;
	for _, start :=  range (startArray) {
		index, res := 0, 0
		temp := start
		for temp[2] != 'Z' {
			if (navigateSequence[index] == 'R') {
				temp = nodeMap[temp][1]
			} else {
				temp = nodeMap[temp][0]
			}
			index = (index + 1) % len(navigateSequence)
			res += 1
		}
		timeToZ = append(timeToZ, res)
	}
	
	file.Close()
	fmt.Println(lcm(timeToZ))
}