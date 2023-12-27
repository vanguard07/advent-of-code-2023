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

func main ()  {
	file, err := os.Open("day09/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		numStrings := strings.Fields(line)
		var nums []int
		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		var lastNums []int
		lastNums = append(lastNums, nums[len(nums) - 1])

		zeroCount := 0
		for (len(nums) != zeroCount) {
			var temp []int
			zeroCount = 0
			for i := 0; i < len(nums) - 1; i++ {
				diff := nums[i + 1] - nums[i]
				if (diff == 0) {
					zeroCount += 1
				}
				temp = append(temp, diff)
			}
			nums = temp
			lastNums = append(lastNums, nums[len(nums) - 1])
		}
		res := 0
		for i := len(lastNums) - 1; i >= 0; i-- {
			res += lastNums[i]
		}
		ans += res
	}
	file.Close()
	fmt.Println(ans)
}