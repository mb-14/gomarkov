package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
	"github.com/montanaflynn/stats"
)

const minimumProbability = 0.05

func main() {
	chain := buildChain()
	scores := getScores(chain)
	stdDev, _ := stats.StandardDeviation(scores)
	mean, _ := stats.Mean(scores)
	for _, data := range getDataset("test.txt") {
		score := sequenceProbablity(chain, data)
		normalizedScore := (score - mean) / stdDev
		isGibberish := normalizedScore < 0
		fmt.Printf("%s | Score: %f | Gibberish: %t\n", data, normalizedScore, isGibberish)
	}
}

func buildChain() *gomarkov.Chain {
	chain := gomarkov.NewChain(3)
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
