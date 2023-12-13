package main

import (
    "bufio"
    "fmt"
    "os"
	"log"
	"unicode"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main()  {
	file, err := os.Open("input.txt")
    check(err)

	scanner := bufio.NewScanner(file)
	ans := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())

		start := 0
		end := len(line) - 1
		var firstNum, lastNum int = 0, 0
		for start <= end {
			if (!unicode.IsDigit(line[start])) {
				start += 1
			}

			if (!unicode.IsDigit(line[end])) {
				end -= 1
			}

			if (unicode.IsDigit(line[start]) && unicode.IsDigit(line[end])) {
				firstNum = int(line[start]) - 48
				lastNum = int(line[end]) - 48
				break
			}
		}
		ans += firstNum * 10 + lastNum

    }

	file.Close()
	fmt.Println(ans)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}