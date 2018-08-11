package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
	"github.com/montanaflynn/stats"
)

const minimumProbability = 0.05

type model struct {
	Mean   float64         `json:"mean"`
	StdDev float64         `json:"std_dev"`
	Chain  *gomarkov.Chain `json:"chain"`
}

func main() {
	train := flag.Bool("train", false, "Train the markov chain")
	username := flag.String("u", "", "Username to classify")
	flag.Parse()
	if *train {
		model := buildModel()
		saveModel(model)
	} else {
		if len(*username) == 0 {
			flag.Usage()
			return
		}
		model, err := loadModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		score := sequenceProbablity(model.Chain, *username)
		normalizedScore := (score -  model.Mean) / model.StdDev
		isGibberish := normalizedScore < 0
		fmt.Printf("Score: %f | Gibberish: %t\n", normalizedScore, isGibberish)
	}
}

func buildModel() model {
	var model model
	chain := buildChain()
	scores := getScores(chain)
	model.StdDev, _ = stats.StandardDeviation(scores)
	model.Mean, _ = stats.Mean(scores)
	model.Chain = chain
	return model
}

func saveModel(model model) {
	jsonObj, _ := json.Marshal(model)
	err := ioutil.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func loadModel() (model, error) {
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		return model{}, err
	}
	var m model
	err = json.Unmarshal(data, &m)
	if err != nil {
		return model{}, err
	}
	return m, nil
}

func buildChain() *gomarkov.Chain {
	chain := gomarkov.NewChain(2)
	for _, data := range getDataset("usernames.txt") {
		chain.Add(split(data))
	}
	return chain
}

func getScores(chain *gomarkov.Chain) []float64 {
	scores := make([]float64, 0)
	for _, data := range getDataset("train.txt") {
		score := sequenceProbablity(chain, data)
		scores = append(scores, score)
	}
	return scores
}

func getDataset(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

func split(str string) []string {
	return strings.Split(str, "")
}

func sequenceProbablity(chain *gomarkov.Chain, input string) float64 {
	tokens := split(input)
	logProb := float64(0)
	pairs := gomarkov.MakePairs(tokens, chain.Order)
	for _, pair := range pairs {
		prob, _ := chain.TransitionProbability(pair.NextState, pair.CurrentState)
		if prob > 0 {
			logProb += math.Log10(prob)
		} else {
			logProb += math.Log10(minimumProbability)
		}
	}
	return math.Pow(10, logProb/float64(len(pairs)))
}
