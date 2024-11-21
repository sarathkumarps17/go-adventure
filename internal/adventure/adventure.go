package adventure

import (
	"fmt"

	"strconv"

	"math/rand"
)

type Option struct {
	Text string
	Arc  string
}
type Chapter struct {
	Title   string
	Story   []string
	Options []Option
}
type Adventure map[string]Chapter

type AdventureInterface interface {
	StartAdventure() Chapter
	GetChapterByName(string) Chapter
	GetChapterList() []string
}

type ChapterInterface interface {
	GetTitle() string
	GetStory() []string
	GetOptions() []Option
	GetArcFromOptionNumber(string) (string error)
	ReadStory(string)
}

// GetChapterList returns a list of all chapter names in the Adventure.
// It traverses the Adventure map and collects the keys (chapter names)
// into a slice of strings.
func (a *Adventure) GetChapterList() []string {
	chapters := make([]string, 0)
	for name := range *a {
		chapters = append(chapters, name)
	}
	return chapters
}

// GetChapterByName takes a chapter name and returns the associated Chapter
// from the Adventure. If the chapter name is not found in the Adventure,
// it returns an empty Chapter.
func (a *Adventure) GetChapterByName(name string) Chapter {
	return (*a)[name]
}

// StartAdventure selects a random chapter from the available chapters
// and returns it. It utilizes the GetChapterList method to retrieve
// the list of chapters and GetChapterByName to fetch the chapter
// by its name.
func (a *Adventure) StartAdventure() Chapter {
	chapterList := a.GetChapterList()
	index := rand.Intn(len(chapterList))
	return a.GetChapterByName(chapterList[index])
}

// GetTitle returns the title of the chapter.
func (c *Chapter) GetTitle() string {
	return c.Title
}

func (c *Chapter) GetStory() []string {
	return c.Story
}

func (c *Chapter) GetOptions() []Option {
	return c.Options
}

func (c *Chapter) ReadStory() {
	fmt.Println(c.Title)
	for _, story := range c.Story {
		fmt.Println(story)
	}
}

// GetArcFromOptionNumber takes an option number as a string, parses it,
// and returns the arc associated with that option in the chapter.
// If the option number is invalid or out of range, it returns an error.
func (c *Chapter) GetArcFromOptionNumber(optionNumber string) (string, error) {
	num, err := strconv.ParseInt(optionNumber, 10, 64)
	if err != nil {
		return "", err
	}
	if num <= 0 || num > int64(len(c.Options)) {
		return "", fmt.Errorf("invalid option number")
	}
	return c.Options[num-1].Arc, nil
}
