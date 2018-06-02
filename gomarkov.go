package gomarkov

import (
	"fmt"
	"math"
	"strings"
)

type list []string

func (l list) key() string {
	return strings.Join(l, "_")
}

//Chain is a markov chain instance
type Chain struct {
	Order  int
	states map[string]state
}

//NewChain creates an instance of Chain
func NewChain(order int) Chain {
	chain := Chain{Order: order}
	chain.states = make(map[string]state)
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

//Learn ...
func (chain *Chain) Learn(input []string) {
	startToken := fill("^", chain.Order-1)
	endToken := fill("$", chain.Order-1)
	tokens := make([]string, 0)
	tokens = append(tokens, startToken...)
	tokens = append(tokens, input...)
	tokens = append(tokens, endToken...)
	nGrams := getNGrams(tokens, chain.Order)
	for i := 0; i < len(nGrams); i++ {
		currentStateIndex, nextNodeIndex := splitNGram(nGrams[i])
		currentState, ok := chain.states[currentStateIndex.key()]
		if !ok {
			currentState = state{
				sum:   0,
				Nodes: make(map[string]node, 0),
			}
		}
		nextNode, ok := currentState.Nodes[nextNodeIndex]
		if !ok {
			nextNode = node{Count: 0}
		}
		nextNode.Count++
		currentState.sum++
		currentState.Nodes[nextNodeIndex] = nextNode
		chain.states[currentStateIndex.key()] = currentState
	}

	for key, state := range chain.states {
		for key, nextNode := range state.Nodes {
			probability := float64(nextNode.Count) / float64(state.sum)
			nextNode.TransitionProbability = math.Log10(probability)
			state.Nodes[key] = nextNode
		}
		chain.states[key] = state
	}
	fmt.Println(chain.states)
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

//Predict returns the joint probability of a text seqeunce based on the transitional probability map
func (chain *Chain) Predict(text []string) float64 {
	logProb := 0.0
	nGrams := getNGrams(text, chain.Order)
	for i := 0; i < len(nGrams); i++ {
		currentStateIndex, nextNodeIndex := splitNGram(nGrams[i])
		if currentState, ok := chain.states[currentStateIndex.key()]; ok {
			if nextNode, ok := currentState.Nodes[nextNodeIndex]; ok {
				logProb += nextNode.TransitionProbability
			}
		}
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
