package adventure

import (
	"reflect"
	"testing"

	"github.com/sarathkumar17/go-adventure/utils"
)

var adventure = Adventure{
	"Chapter 1": {
		Title: "Chapter 1",
		Story: []string{"Story 1", "Story 2"},
		Options: []Option{
			{
				Text: "Option 1",
				Arc:  "Arc 1",
			},
			{
				Text: "Option 2",
				Arc:  "Arc 2",
			},
		},
	},
	"Chapter 2": {
		Title: "Chapter 2",
		Story: []string{"Story 1", "Story 2"},
		Options: []Option{
			{
				Text: "Option 1",
				Arc:  "Chapter 1",
			},
			{
				Text: "Option 2",
				Arc:  "Chapter 3",
			},
		},
	},
	"Chapter 3": {
		Title: "Chapter 3",
		Story: []string{"Story 1", "Story 2"},
		Options: []Option{
			{
				Text: "Option 1",
				Arc:  "Chapter 2",
			},
			{
				Text: "Option 2",
				Arc:  "Chapter 1",
			},
		},
	},
}

// Returns a list of chapter names when Adventure has multiple chapters
func TestGetChapterListMultipleChapters(t *testing.T) {

	chapters := []string{"Chapter 1", "Chapter 2", "Chapter 3"}
	result := adventure.GetChapterList()
	// compare the chapters in result with the expected list of chapters without the order
	if !utils.CheckSlicesHaveSameElements(chapters, result) {
		t.Errorf("Expected %v, but got %v", chapters, result)
	} //CheckSlicesHaveSameElements(result, chapters)

}

// Returns an empty list when Adventure has no chapters
func TestGetChapterListNoChapters(t *testing.T) {
	adventure := Adventure{}
	expected := []string{}
	result := adventure.GetChapterList()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetChapterByName(t *testing.T) {
	expected := "Chapter 1"
	result := adventure.GetChapterByName("Chapter 1").Title
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestStartAdventure(t *testing.T) {
	chapters := adventure.GetChapterList()
	result := adventure.StartAdventure().Title
	isRandomChapter := false
	for _, chapter := range chapters {
		if chapter == result {
			isRandomChapter = true
		}
	}

	if !isRandomChapter {
		t.Errorf("Expected a random chapter, but got %v", result)
	}
}

// Retrieve the title of a chapter successfully
func TestGetTitleSuccess(t *testing.T) {
	chapter := &Chapter{Title: "The Adventure Begins"}
	title := chapter.GetTitle()
	if title != "The Adventure Begins" {
		t.Errorf("expected 'The Adventure Begins', got '%s'", title)
	}
}

// Handle cases where the chapter title is an empty string
func TestGetTitleEmptyString(t *testing.T) {
	chapter := &Chapter{Title: ""}
	title := chapter.GetTitle()
	if title != "" {
		t.Errorf("expected empty string, got '%s'", title)
	}
}

func TestGetStorySuccess(t *testing.T) {
	chapter := &Chapter{Story: []string{"Story 1", "Story 2"}}
	story := chapter.GetStory()
	if !reflect.DeepEqual(story, []string{"Story 1", "Story 2"}) {
		t.Errorf("expected ['Story 1', 'Story 2'], got %v", story)
	}
}

func TestGetOptionsSuccess(t *testing.T) {
	chapter := &Chapter{Options: []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}}
	options := chapter.GetOptions()
	if !reflect.DeepEqual(options, []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}) {
		t.Errorf("expected [{Text: 'Option 1', Arc: 'Arc 1'}, {Text: 'Option 2', Arc: 'Arc 2'}], got %v", options)
	}
}

func TestReadStorySuccess(t *testing.T) {
	chapter := &Chapter{Story: []string{"Story 1", "Story 2"}}
	chapter.ReadStory()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected no panic, got %v", r)
		}
	}()

}

func TestGetArcFromOptionNumberSuccess(t *testing.T) {
	chapter := &Chapter{Options: []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}}
	arc, err := chapter.GetArcFromOptionNumber("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if arc != "Arc 1" {
		t.Errorf("expected 'Arc 1', got %s", arc)
	}
}

func TestGetArcFromOptionNumberInvalidOptionNumber(t *testing.T) {
	chapter := &Chapter{Options: []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}}
	_, err := chapter.GetArcFromOptionNumber("3")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetArcFromOptionNumberOutOfRange(t *testing.T) {
	chapter := &Chapter{Options: []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}}
	_, err := chapter.GetArcFromOptionNumber("0")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetArcFromOptionNumberEmptyOptionNumber(t *testing.T) {
	chapter := &Chapter{Options: []Option{{Text: "Option 1", Arc: "Arc 1"}, {Text: "Option 2", Arc: "Arc 2"}}}
	_, err := chapter.GetArcFromOptionNumber("")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
