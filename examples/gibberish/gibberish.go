package main

import (
	"math"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
)

func main() {
	goodChain := gomarkov.NewChain(4)
	goodTrain, goodTest := getData("good.txt", 50)
	badTrain, badTest := getData("bad.txt", 10)

	for _, data := range goodTrain {
		goodChain.Add(split(data))
	}

	for _, data := range append(goodTest) {
		match := goodChain.Match(split(data))
		y := 1/(1 + math.Exp(-match))
		fmt.Printf("%g, %g\n", match, y)
	}

	// badChain := gomarkov.NewChain(3)

	// for _, data := range badTrain {
	// 	badChain.Add(split(data))
	// }

	for _, data := range append(badTrain, badTest...) {
		match := goodChain.Match(split(data))
		y := 1/(1 + math.Exp(-match))
		fmt.Printf("%g, %g\n", match, y)
	}
}

func getData(fileName string, n int) ([]string, []string) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	list = shuffle(list)
	return list[n:], list[len(list)-n:]
}

func shuffle(arr []string) []string {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func split(str string) []string {
	return strings.Split(str, "")
}
