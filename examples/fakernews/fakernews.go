package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/mb-14/gomarkov"
)

const (
	hnBaseURL        = "https://hacker-news.firebaseio.com/v0/"
	hnTopStoriesPath = "topstories.json"
	hnStoryItemPath  = "item/"
)

type hnStory struct {
	Title string `json:"title"`
}

func main() {
	train := flag.Bool("train", false, "Train the markov chain")
	flag.Parse()
	if *train {
		chain, err := buildModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		saveModel(chain)
	} else {
		chain, err := loadModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		generateHNStory(chain)
	}
}

func buildModel() (*gomarkov.Chain, error) {
	stories, err := fetchHNTopStories()
	if err != nil {
		return nil, err
	}
	chain := gomarkov.NewChain(1)
	var wg sync.WaitGroup
	wg.Add(len(stories))
	fmt.Println("Adding HN story titles to markov chain...")
	for _, storyID := range stories {
		go func(storyID int) {
			defer wg.Done()
			story, err := fetchHNStory(storyID)
			if err != nil {
				fmt.Println(err)
				return
			}
			chain.Add(strings.Split(story.Title, " "))
		}(storyID)
	}
	wg.Wait()
	return chain, nil
}

func loadModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func fetchHNTopStories() ([]int, error) {
	fmt.Println("Fetching HN top stories...")
	resp, err := http.Get(fmt.Sprintf("%s%s", hnBaseURL, hnTopStoriesPath))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var stories []int
	err = json.Unmarshal(body, &stories)
	return stories, err
}

func fetchHNStory(storyID int) (hnStory, error) {
	var story hnStory
	resp, err := http.Get(fmt.Sprintf("%s%s%d.json", hnBaseURL, hnStoryItemPath, storyID))
	if err != nil {
		return story, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return story, err
	}
	err = json.Unmarshal(body, &story)
	return story, err
}

func saveModel(chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func generateHNStory(chain *gomarkov.Chain) {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
}
