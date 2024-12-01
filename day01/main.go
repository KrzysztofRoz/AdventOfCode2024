package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LocationsOnList struct {
	LocationID int
	Repetition int
}

func main() {
	filePath := "/home/krzysztofroz/Repos/AdventOfCode2024/inputs/input01.txt"
	fmt.Println("=======================Task1=======================")
	ans1 := FirstTaskDay01(filePath)
	fmt.Println("First task:")
	fmt.Println(ans1)
	fmt.Println("=======================Task2=======================")
	ans2 := SecondTaskDay01(filePath)
	fmt.Println("Second task:")
	fmt.Println(ans2)

}
func FirstTaskDay01(path string) int {
	right, left, err := ParseInputToSlices(path)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	if len(right) != len(left) {
		log.Fatal("The slices len do not match")
	}
	distSum := 0
	for i := range right {
		distSum += Distance(right[i], left[i])
	}

	return distSum
}

func SecondTaskDay01(path string) int {
	right, left, err := ParseInputToSlices(path)
	if err != nil {
		log.Fatal(err)
	}
	if len(right) != len(left) {
		log.Fatal("The slices len do not match")
	}
	rightAggregated := CountRepetition(right)
	leftAggregated := CountRepetition(left)
	similarityScore := 0
	for loc, count := range leftAggregated {
		rep, exist := rightAggregated[loc]
		if !exist {
			similarityScore += 0
		} else {
			similarityScore += loc * rep * count
		}
	}

	return similarityScore
}
func ParseInputToSlices(filePath string) (rightSlice, leftSlice []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rightSlice = []int{}
	leftSlice = []int{}
	taskSeparator := "   "

	for scanner.Scan() {
		slicesValues := strings.Split(scanner.Text(), taskSeparator)

		rightValue, err := strconv.Atoi(slicesValues[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		leftValue, err := strconv.Atoi(slicesValues[0])
		if err != nil {
			return []int{}, []int{}, err
		}

		leftSlice = append(leftSlice, leftValue)
		rightSlice = append(rightSlice, rightValue)
	}

	return rightSlice, leftSlice, nil
}

func Distance(a, b int) int {
	result := a - b
	if result < 0 {
		return -result
	}
	return result
}
func CountRepetition(slice []int) map[int]int {
	locations := make(map[int]int)
	for _, s := range slice {
		_, exist := locations[s]
		if exist {
			locations[s]++
		} else {
			locations[s] = 1
		}
	}
	return locations
}
