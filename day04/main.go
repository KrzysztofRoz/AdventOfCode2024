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
	// fmt.Println("=======================Task2=======================")
	// ans2 := SecondTaskDay04(filePath)
	// fmt.Println("Second task:")
	// fmt.Println(ans2)

}

func FirstTaskDay04(path string) int {
	text, xpos, err := ParseInput(path)
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

func ParseInput(filePath string) ([][]string, []Position, error) {

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
		count += CheckWord(xpos, word, text)
	}

	return count, err
}

func CheckWord(xpos Position, word string, text [][]string) int {
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
