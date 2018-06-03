package gomarkov

import (
	"math"
	"strings"
	"sync"
)

const minimumProbability = 0.1

type list []string

func (l list) key() string {
	return strings.Join(l, "_")
}

//Chain is a markov chain instance
type Chain struct {
	Order        int
	statePool    *spool
	frequencyMat map[int]sparseArray
	sync.RWMutex
}

type sparseArray map[int]int

func (s sparseArray) sum() int {
	sum := 0
	for _, count := range s {
		sum += count
	}
	return sum
}

//NewChain creates an instance of Chain
func NewChain(order int) *Chain {
	chain := Chain{Order: order}
	chain.statePool = &spool{stringMap: make(map[string]int)}
	chain.frequencyMat = make(map[int]sparseArray, 0)
	return &chain
}

//Add ...
func (chain *Chain) Add(input []string) {
	startToken := fill("^", chain.Order-1)
	endToken := fill("$", chain.Order-1)
	tokens := make([]string, 0)
	tokens = append(tokens, startToken...)
	tokens = append(tokens, input...)
	tokens = append(tokens, endToken...)
	nGrams := getNGrams(tokens, chain.Order)
	for i := 0; i < len(nGrams); i++ {
		current, next := splitNGram(nGrams[i])
		currentIndex := chain.statePool.add(current.key())
		nextIndex := chain.statePool.add(next)
		chain.Lock()
		if chain.frequencyMat[currentIndex] == nil {
			chain.frequencyMat[currentIndex] = make(sparseArray, 0)
		}
		chain.frequencyMat[currentIndex][nextIndex]++
		chain.Unlock()
	}
}

func splitNGram(ngram []string) (list, string) {
	return ngram[:len(ngram)-1], ngram[len(ngram)-1]
}

func fill(value string, count int) []string {
	arr := make([]string, count)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

//Match returns the joint probability of a text seqeunce based on the transitional probability map
func (chain *Chain) Match(text []string) float64 {
	logProb := 0.0
	nGrams := getNGrams(text, chain.Order)
	for i := 0; i < len(nGrams); i++ {
		current, next := splitNGram(nGrams[i])
		if currentIndex, ok := chain.statePool.get(current.key()); ok {
			if nextIndex, ok := chain.statePool.get(next); ok {
				arr := chain.frequencyMat[currentIndex]
				sum := arr.sum()
				count := arr[nextIndex]
				if sum > 0 && count > 0 {
					prob := float64(count) / float64(sum)
					logProb += math.Log10(prob)
					continue
				}
			}
		}
		logProb += math.Log10(minimumProbability)
	}
	return math.Pow(10, logProb/float64(len(nGrams)))
}

func getNGrams(tokens []string, order int) [][]string {
	var nGrams [][]string
	for i := 0; i < len(tokens)-order+1; i++ {
		nGrams = append(nGrams, tokens[i:i+order])
	}
	return nGrams
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
