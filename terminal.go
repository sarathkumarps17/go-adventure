package main

import (
	"fmt"

	"github.com/sarathkumar17/go-adventure/internal/adventure"

	jsonparser "github.com/sarathkumar17/go-adventure/pkg/json_parser"
)

func main() {
	fmt.Println("Welcome to the adventure!")

	adventure, err := readAdventure("gopher.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	firstChapter := adventure.StartAdventure()
	firstChapter.ReadStory()

	for {
		askOption(firstChapter)
		optionText := readOptionNumber()
		if optionText == "yes" {
			firstChapter = adventure.StartAdventure()
			firstChapter.ReadStory()
		}
		if optionText == "no" {
			fmt.Printf("Bye!\n")
			break
		}
		arc, err := firstChapter.GetArcFromOptionNumber(optionText)
		if err != nil {
			fmt.Printf("Invalid Option: %v\n", err)
			continue
		}
		fmt.Printf("Moving to chapter: %s\n", arc)

		firstChapter = adventure.GetChapterByName(arc)
		firstChapter.ReadStory()
	}
}

func readAdventure(filename string) (adventure.Adventure, error) {
	bs, err := jsonparser.ReadJSON(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return jsonparser.ParseJSON(bs)
}

func askOption(chapter adventure.Chapter) {
	fmt.Println("What would you like to do?")
	options := chapter.GetOptions()
	if len(options) == 0 {
		fmt.Println("You are at the end of the story, would you like to play again? (yes/no)")
		return
	}
	fmt.Printf("Please enter your option number: \n")
	for i, option := range options {
		fmt.Printf("%d - %s\n", i+1, option.Text)
	}
}

func readOptionNumber() string {
	var optionNumber string
	fmt.Scan(&optionNumber)
	return optionNumber
}
