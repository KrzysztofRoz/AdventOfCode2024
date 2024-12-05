package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

func (p Position) MoveByVector(vector Position) (newPosition *Position) {
	return &Position{
		X: p.X + vector.X,
		Y: p.Y + vector.Y,
	}
}
func (p Position) GetLetter(text [][]string) string {

	return text[p.Y][p.X]
}

func (p Position) GetWord(length int, vector Position, text [][]string) string {
	var builder strings.Builder

	newPos := p
	for i := 0; i < length; i++ {
		newPos = *newPos.MoveByVector(vector)
		builder.WriteString(newPos.GetLetter(text))
	}
	return builder.String()
}

func main() {
	filePath := "/home/krzysztofroz/Repos/AdventOfCode2024/inputs/input04.txt"
	fmt.Println("=======================Task1=======================")
	ans1 := FirstTaskDay04(filePath)
	fmt.Println("First task:")
	fmt.Println(ans1)
	fmt.Println("=======================Task2=======================")
	ans2 := SecondTaskDay04(filePath)
	fmt.Println("Second task:")
	fmt.Println(ans2)

}

// First part solution
func FirstTaskDay04(path string) int {
	text, xpos, err := ParseInputForXMAS(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count, err := SearchXMAS(xpos, text)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return count
}

func ParseInputForXMAS(filePath string) ([][]string, []Position, error) {

	textMatrix := make([][]string, 0)
	xPosition := []Position{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return textMatrix, xPosition, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		textMatrix = append(textMatrix, letters)
		for i, letter := range letters {
			if letter == "X" {
				xPosition = append(xPosition, Position{X: i, Y: count}) //Get all X coordinates for faster word search
			}
		}
		count++
	}
	return textMatrix, xPosition, nil
}
func SearchXMAS(xs []Position, text [][]string) (count int, err error) {
	word := "XMAS"

	count = 0
	for _, xpos := range xs {
		count += CheckXMASWord(xpos, word, text)
	}

	return count, err
}

func CheckXMASWord(xpos Position, word string, text [][]string) int {
	hight := len(text)
	width := len(text[0])
	wordLength := len(word) - 1 //sub one beacouse we already got one letter X
	count := 0
	moveVector := Position{}

	//check up word backwards
	if xpos.Y-wordLength >= 0 {
		moveVector = Position{X: 0, Y: -1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	// check up right diagonal word
	if xpos.Y-wordLength >= 0 && xpos.X+wordLength < width {
		moveVector = Position{X: 1, Y: -1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}

	}

	//check right word
	if xpos.X+wordLength < width {
		moveVector = Position{X: 1, Y: 0}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	//check bottom right diagonal word
	if xpos.X+wordLength < width && xpos.Y+wordLength < hight {
		moveVector = Position{X: 1, Y: 1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	//check down word
	if xpos.Y+wordLength < hight {
		moveVector = Position{X: 0, Y: 1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	//check down left diagonal word(backwards)
	if xpos.Y+wordLength < hight && xpos.X-wordLength >= 0 {
		moveVector = Position{X: -1, Y: 1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	//check left word (backwards)
	if xpos.X-wordLength >= 0 {
		moveVector = Position{X: -1, Y: 0}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	//check up left word (backwards)
	if xpos.X-wordLength >= 0 && xpos.Y-wordLength >= 0 {
		moveVector = Position{X: -1, Y: -1}
		checkWord := fmt.Sprintf("X%s", xpos.GetWord(wordLength, moveVector, text))
		if checkWord == word {
			count++
		}
	}

	return count
}

// Solution for second part

func SecondTaskDay04(path string) int {
	text, apos, err := ParseInputForMAS(path)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count, err := SearchMAS(apos, text)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return count
}

func ParseInputForMAS(filePath string) ([][]string, []Position, error) {

	textMatrix := make([][]string, 0)
	aPosition := []Position{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return textMatrix, aPosition, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		textMatrix = append(textMatrix, letters)
		for i, letter := range letters {
			if letter == "A" {
				aPosition = append(aPosition, Position{X: i, Y: count}) //Get all A coordinates for faster word search
			}
		}
		count++
	}
	return textMatrix, aPosition, nil
}
func SearchMAS(as []Position, text [][]string) (count int, err error) {
	word := "MAS"

	count = 0
	for _, apos := range as {
		count += CheckMASWord(apos, word, text)
	}

	return count, err
}

func CheckMASWord(apos Position, word string, text [][]string) int {
	hight := len(text)
	width := len(text[0])

	count := 0

	if apos.X-1 >= 0 && apos.X+1 < width && apos.Y-1 >= 0 && apos.Y+1 < hight {
		var diagWord1 string
		var diagWord2 string
		var builder strings.Builder

		// from up to down word
		builder.WriteString(Position{Y: apos.Y - 1, X: apos.X - 1}.GetLetter(text))
		builder.WriteString(apos.GetLetter(text))
		builder.WriteString(Position{Y: apos.Y + 1, X: apos.X + 1}.GetLetter(text))

		diagWord1 = builder.String()

		if diagWord1 == "MAS" || reverse(diagWord1) == "MAS" {
			// from down to up word
			builder = strings.Builder{}
			builder.WriteString(Position{Y: apos.Y + 1, X: apos.X - 1}.GetLetter(text))
			builder.WriteString(apos.GetLetter(text))
			builder.WriteString(Position{Y: apos.Y - 1, X: apos.X + 1}.GetLetter(text))
			diagWord2 = builder.String()
			if diagWord2 == "MAS" || reverse(diagWord2) == "MAS" {
				count++
			}
		}
	}

	return count
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
