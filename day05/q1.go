package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Pair struct {
	Key   int
	Value int
}

func getNumber(s string) int {
	sanitizedString := strings.TrimSpace(s)
	x, err := strconv.ParseInt(sanitizedString, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(x)
}

func createTuple(scanner *bufio.Scanner) [][]int {
	var tuples [][]int
	for scanner.Scan() {
		line := scanner.Text()
		listing := strings.Split(line, " ")
		if line != "" {
			var destination, source, rge int = getNumber(listing[0]), getNumber(listing[1]), getNumber(listing[2])
			entry := []int{destination, source, rge}
			tuples = append(tuples, entry)
		} else {
			break
		}
	}
	return tuples
}

func getMapping(tuples [][]int, input int) int {
	for _, tuple := range tuples {
		if input >= tuple[1] && input < tuple[1]+tuple[2] {
			return tuple[0] + (input - tuple[1])
		}
	}
	return input
}

func main() {
	file, err := os.Open("day05/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	ans := math.MaxInt
	var seedLine string
	if scanner.Scan() {
		seedLine = scanner.Text()
	}

	numStrings := strings.Fields(strings.Split(seedLine, ":")[1])
	var seeds []int

	for _, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		seeds = append(seeds, num)
	}

	var seedSoil [][]int
	var soilFert [][]int
	var fertWater [][]int
	var waterLight [][]int
	var lightTemp [][]int
	var tempHumid [][]int
	var humidLoc [][]int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ":")[0]
		switch line {
		case "seed-to-soil map":
			seedSoil = createTuple(scanner)
		case "soil-to-fertilizer map":
			soilFert = createTuple(scanner)
		case "fertilizer-to-water map":
			fertWater = createTuple(scanner)
		case "water-to-light map":
			waterLight = createTuple(scanner)
		case "light-to-temperature map":
			lightTemp = createTuple(scanner)
		case "temperature-to-humidity map":
			tempHumid = createTuple(scanner)
		case "humidity-to-location map":
			humidLoc = createTuple(scanner)
		}
	}

	for _, seed := range seeds {
		soil := getMapping(seedSoil, seed)
		fertilizer := getMapping(soilFert, soil)
		water := getMapping(fertWater, fertilizer)
		light := getMapping(waterLight, water)
		temperature := getMapping(lightTemp, light)
		humidity := getMapping(tempHumid, temperature)
		location := getMapping(humidLoc, humidity)
		fmt.Println(seed, soil, fertilizer, water, light, temperature, humidity, location)
		if location < ans {
			ans = location
		}
	}

	file.Close()
	fmt.Println(ans)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
