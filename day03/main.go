package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Operation struct {
	Instruction string
	X           int
	Y           int
}

func main() {
	filePath := "/home/krzysztofroz/Repos/AdventOfCode2024/inputs/input03.txt"
	fmt.Println("=======================Task1=======================")
	ans1 := FirstTaskDay03(filePath)
	fmt.Println("First task:")
	fmt.Println(ans1)
	fmt.Println("=======================Task2=======================")
	ans2 := SecondTaskDay03(filePath)
	fmt.Println("Second task:")
	fmt.Println(ans2)

}

func FirstTaskDay03(path string) int {
	operations, err := ParseInput(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count := 0
	for _, ops := range operations {
		if ops.Instruction == "mul" {
			count += ops.X * ops.Y
		}
	}
	return count
}

func ParseInput(filePath string) (operations []Operation, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		muls, err := ParseLineToOperations(scanner.Text())
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		operations = append(operations, muls...)
	}
	return operations, nil
}

func ParseLineToOperations(line string) (operations []Operation, err error) {
	r := regexp.MustCompile(`(mul\([0-9]{1,3}\,[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	matchMul := r.FindAllString(line, -1)

	numberRegex := regexp.MustCompile(`[0-9]{1,3}`)
	doRegex := regexp.MustCompile(`do\(\)`)
	for _, match := range matchMul {
		numbers := numberRegex.FindAllString(match, -1)
		if len(numbers) > 0 {
			x, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatal(err)
				return operations, err
			}
			y, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatal(err)
				return operations, err
			}
			operations = append(operations, Operation{X: x, Y: y, Instruction: "mul"})
		} else if doRegex.MatchString(match) {
			operations = append(operations, Operation{Instruction: "do"})
		} else {
			operations = append(operations, Operation{Instruction: "dont"})
		}
	}
	return operations, nil
}

func SecondTaskDay03(path string) int {
	operations, err := ParseInput(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count := 0
	IsEnabled := true
	for _, ops := range operations {
		switch ops.Instruction {
		case "do":
			IsEnabled = true
		case "dont":
			IsEnabled = false
		case "mul":
			if IsEnabled {
				count += ops.X * ops.Y
			}
		}
	}
	return count
}
