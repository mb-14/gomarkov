package gomarkov

import "strings"

//Pair is a pair of consecutive states in a sequece
type Pair struct {
	CurrentState NGram  // n = order of the chain
	NextState    string // n = 1
}

//NGram is a array of words
type NGram []string

type sparseArray map[int]int

func (ngram NGram) key() string {
	return strings.Join(ngram, "_")
}

func (s sparseArray) sum() int {
	sum := 0
	for _, count := range s {
		sum += count
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func array(value string, count int) []string {
	arr := make([]string, count)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

//MakePairs generates n-gram pairs of consecutive states in a sequence
func MakePairs(tokens []string, order int) []Pair {
	var pairs []Pair
	for i := 0; i < len(tokens)-order; i++ {
		pair := Pair{
			CurrentState: tokens[i : i+order],
			NextState:    tokens[i+order],
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
