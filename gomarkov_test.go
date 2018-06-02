package gomarkov

import (
	"fmt"
	"strings"
	"testing"
)

func TestChain(t *testing.T) {
	chain := NewChain(2)
	chain.Learn(strings.Split("The grass is greener in China", " "))
	chain.Learn(strings.Split("The grass is greener in America", " "))
	prediction := chain.Predict(strings.Split("The grass is greener in America", " "))
	fmt.Println(prediction)
	t.Fail()
}
