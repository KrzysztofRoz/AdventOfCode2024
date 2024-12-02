package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Content      []int
	IsIncreasing bool
	IsDecreasing bool
	IsSafe       bool
}

func main() {
	filePath := "/home/krzysztofroz/Repos/AdventOfCode2024/inputs/input02.txt"
	fmt.Println("=======================Task1=======================")
	ans1 := FirstTaskDay02(filePath)
	fmt.Println("First task:")
	fmt.Println(ans1)
	// fmt.Println("=======================Task2=======================")
	// ans2 := SecondTaskDay02(filePath)
	// fmt.Println("Second task:")
	// fmt.Println(ans2)

}

func FirstTaskDay02(path string) int {
	repots, err := ParseInputToReports(path)
	if err != nil {
		panic(err)
	}
	count := 0

	for _, rep := range repots {
		if rep.IsSafe {
			count++
		}
	}

	return count
}

func ParseInputToReports(filePath string) (reports []Report, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reportSeparator := " "
	for scanner.Scan() {
		reportData := strings.Split(scanner.Text(), reportSeparator)
		reportDataAsInt := []int{}
		for _, level := range reportData {
			data, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal("[ParseInputToReports] Parse Daata error")
				return reports, err
			}
			reportDataAsInt = append(reportDataAsInt, data)
		}
		rep := Report{Content: reportDataAsInt}
		rep.CheckIfSafe()
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *Report) CheckIfSafe() {
	if len(r.Content) < 1 {
		fmt.Println(r)
		fmt.Println("Report empty")
		return
	}
	if len(r.Content) < 2 {
		fmt.Println("Only one record")
		r.IsSafe = true
		return
	}
	maxChange := 3
	var checkingIncrease bool
	if (r.Content[0] - r.Content[1]) > 0 {
		checkingIncrease = false
	}
	if (r.Content[0] - r.Content[1]) < 0 {
		checkingIncrease = true
	}
	if (r.Content[0] - r.Content[1]) == 0 {
		fmt.Println("Report is unsafe, no increase or decrease")
		r.IsSafe = false
		return
	}

	if checkingIncrease {
		for i := 0; i < len(r.Content)-1; i++ {
			dif := r.Content[i+1] - r.Content[i]
			if dif < 1 || dif > maxChange {
				r.IsSafe = false
				return
			}
		}
		r.IsIncreasing = true
		r.IsSafe = true

	} else {
		for i := 0; i < len(r.Content)-1; i++ {
			dif := r.Content[i+1] - r.Content[i]
			if dif < -maxChange || dif > -1 {
				r.IsSafe = false
				return
			}
		}
		r.IsDecreasing = true
		r.IsSafe = true
	}

}
