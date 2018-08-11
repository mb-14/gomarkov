# gomarkov

Go implementation of markov chains for textual data. 

You can find out more about markov chains [here](http://setosa.io/ev/markov-chains/) and [here](https://towardsdatascience.com/introduction-to-markov-chains-50da3645a50d)

## Usage
```
import (
        "github.com/mb-14/gomarkov"
        "fmt"
    )

func main() {
    //Create a chain of order 2
    chain := gomarkov.NewChain(2)

    //Feed in training data
    chain.Add([]string{"I", "want", "a", "cheese", "burger"})
    chain.Add([]string{"I", "want", "a", "chilled", "sprite"})
    chain.Add([]string{"I", "want", "to", "go", "to", "the", "movies"})

    //Get transition probability of a sequence
    prob := chain.TransitionProbability("a", []string{"I", "want"})
    fmt.Println(prob)
    //Output: 0.6666666666666666
}
```
## Examples

- [Gibberish username detector](github.com/mb-14/gomarkov/examples/gibberish)