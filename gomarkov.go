package gomarkov

import (
	"math"
	"strings"
)

const minimumProbability = 0.1

type list []string

func (l list) key() string {
	return strings.Join(l, "_")
}

//Chain is a markov chain instance
type Chain struct {
	Order        int
	stateMap     map[string]int
	frequencyMat [][]int
	sumArray     []int
}

//NewChain creates an instance of Chain
func NewChain(order int) Chain {
	chain := Chain{Order: order}
	chain.stateMap = make(map[string]int, 0)
	chain.frequencyMat = make([][]int, 0)
	return chain
}

type state struct {
	sum   int
	Nodes map[string]node
}

type node struct {
	Count                 int
	TransitionProbability float64
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
		lastIndex := len(chain.stateMap)
		currentState, nextNode := splitNGram(nGrams[i])
		currentStateIndex, ok := chain.stateMap[currentState.key()]
		if !ok {
			currentStateIndex = lastIndex
			chain.stateMap[currentState.key()] = lastIndex
			lastIndex++
		}
		nextNodeIndex, ok := chain.stateMap[nextNode]
		if !ok {
			nextNodeIndex = lastIndex
			chain.stateMap[nextNode] = nextNodeIndex
			lastIndex++
		}

		if lastIndex > len(chain.frequencyMat) {
			for i := len(chain.frequencyMat); i < lastIndex; i++ {
				chain.frequencyMat = append(chain.frequencyMat, make([]int, 0))
			}
		}

		if lastIndex > len(chain.frequencyMat[currentStateIndex]) {
			for i := len(chain.frequencyMat[currentStateIndex]); i < lastIndex; i++ {
				chain.frequencyMat[currentStateIndex] = append(chain.frequencyMat[currentStateIndex], 0)
			}
		}
		chain.frequencyMat[currentStateIndex][nextNodeIndex]++
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
		currentState, nextNode := splitNGram(nGrams[i])
		if currentStateIndex, ok := chain.stateMap[currentState.key()]; ok {
			if nextNodeIndex, ok := chain.stateMap[nextNode]; ok {
				arr := chain.frequencyMat[currentStateIndex]
				if nextNodeIndex < len(arr) {
					sum := sum(arr)
					count := arr[nextNodeIndex]
					if sum > 0 && count > 0 {
						prob := float64(count) / float64(sum)
						logProb += math.Log10(prob)
						continue
					}
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

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
