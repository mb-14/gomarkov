package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
	"github.com/montanaflynn/stats"
)

func main() {
	chain := buildChain()

	scores := getScores(chain)
	stdDev, _ := stats.StandardDeviation(scores)
	mean, _ := stats.Mean(scores)
	for _, data := range getDataset("test.txt") {
		score := chain.Match(split(data))
		normalizedScore := (score - mean) / stdDev
		isGibberish := normalizedScore < 0
		fmt.Printf("%s - Gibberish: %t\n", data, isGibberish)
	}
}

func buildChain() *gomarkov.Chain {
	chain := gomarkov.NewChain(2)
	for _, data := range getDataset("id_list.txt") {
		chain.Add(split(data))
	}
	return chain
}

func getScores(chain *gomarkov.Chain) []float64 {
	scores := make([]float64, 0)
	for _, data := range getDataset("train.txt") {
		score := chain.Match(split(data))
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
