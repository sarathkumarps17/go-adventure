package services

import "github.com/sarathkumar17/go-adventure/internal/adventure"

func GetChapterService(adventure adventure.Adventure, chapterName string) (adventure.Chapter, error) {
	return adventure.GetChapterByName(chapterName), nil
}

func StartStoryService(adventure adventure.Adventure) (adventure.Chapter, error) {
	return adventure.StartAdventure(), nil
}
