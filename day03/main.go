package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type MulObject struct {
	X   int
	Y   int
	Mul int
}

func main() {
	filePath := "/home/krzysztofroz/Repos/AdventOfCode2024/inputs/input03.txt"
	fmt.Println("=======================Task1=======================")
	ans1 := FirstTaskDay03(filePath)
	fmt.Println("First task:")
	fmt.Println(ans1)
	// fmt.Println("=======================Task2=======================")
	// ans2 := SecondTaskDay02(filePath)
	// fmt.Println("Second task:")
	// fmt.Println(ans2)

}

func FirstTaskDay03(path string) int {
	muls, err := ParseInput(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count := 0
	for _, mul := range muls {
		count += mul.Mul
	}
	return count
}

func ParseInput(filePath string) (mulData []MulObject, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		muls, err := ParseLineToMuls(scanner.Text())
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		mulData = append(mulData, muls...)
	}
	return mulData, nil
}

func ParseLineToMuls(line string) (mulData []MulObject, err error) {
	r := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	matchMul := r.FindAllString(line, -1)

	numberRegex := regexp.MustCompile(`[0-9]{1,3}`)
	for _, match := range matchMul {
		numbers := numberRegex.FindAllString(match, -1)

		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
			return mulData, err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
			return mulData, err
		}
		mulData = append(mulData, MulObject{X: x, Y: y, Mul: x * y})

	}
	return mulData, nil
}
