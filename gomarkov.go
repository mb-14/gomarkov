package gomarkov

import (
	"errors"
	"sync"
)

//Chain is a markov chain instance
type Chain struct {
	Order        int
	statePool    *spool
	frequencyMat map[int]sparseArray
	sync.RWMutex
}

//NewChain creates an instance of Chain
func NewChain(order int) *Chain {
	chain := Chain{Order: order}
	chain.statePool = &spool{stringMap: make(map[string]int)}
	chain.frequencyMat = make(map[int]sparseArray, 0)
	return &chain
}

//Add adds the transition counts to the chain for a given sequence
func (chain *Chain) Add(input []string) {
	startToken := fill("^", chain.Order)
	endToken := fill("$", chain.Order)
	tokens := make([]string, 0)
	tokens = append(tokens, startToken...)
	tokens = append(tokens, input...)
	tokens = append(tokens, endToken...)
	pairs := MakePairs(tokens, chain.Order)
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		currentIndex := chain.statePool.add(pair.CurrentState.key())
		nextIndex := chain.statePool.add(pair.NextState)
		chain.Lock()
		if chain.frequencyMat[currentIndex] == nil {
			chain.frequencyMat[currentIndex] = make(sparseArray, 0)
		}
		chain.frequencyMat[currentIndex][nextIndex]++
		chain.Unlock()
	}
}

//TransitionProbability returns the transition probability between two states
func (chain *Chain) TransitionProbability(next string, current NGram) (float64, error) {
	if len(current) != chain.Order {
		return 0, errors.New("N-gram length does not match chain order")
	}
	currentIndex, currentExists := chain.statePool.get(current.key())
	nextIndex, nextExists := chain.statePool.get(next)
	if !currentExists || !nextExists {
		return 0, nil
	}
	arr := chain.frequencyMat[currentIndex]
	sum := float64(arr.sum())
	freq := float64(arr[nextIndex])
	return freq / sum, nil
}
