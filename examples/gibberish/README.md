# Gibberish username detector

## Overview
`gibberish` is a simple program to differentiate between genuine looking and gibberish usernames. It uses a 2nd order markov chain.

## Usage
```
> go build -o gibberish .
> ./gibberish -train
> ./gibberish -u jack.reacher
Score: 0.184727 | Gibberish: false
> ./gibberish -u fgsdgtqh
Score: -0.739162 | Gibberish: true
```

## Explanation

The markov chain trains on a big set of geniune usernames (`usernames.txt`) and records the character transition probabilities.
We then calculate the transition probability scores of a set of genuine looking and gibberish usernames (`train.txt`) and the extract the mean and standard deviation of the scores. We do this to get the z-score of a username's transition probaility score which helps in comparing with the results of the training dataset. A negative z-score indicates a gibberish username and positive one indicates a genuine one.
 