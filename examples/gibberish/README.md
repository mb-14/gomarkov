# Gibberish username detector

## Overview
`gibberish.go` is a simple program to differentiate between genuine looking and gibberish usernames. It uses a 2nd order markov chain.

## Usage
```
> go run gibberish.go
Output:
david.williams | Score: 1.468633 | Gibberish: false
chris.williams | Score: 1.254395 | Gibberish: false
john.williams | Score: 1.529021 | Gibberish: false
andrew.smith | Score: 0.991997 | Gibberish: false
adam.smith | Score: 0.639571 | Gibberish: false
joe.smith | Score: 0.952212 | Gibberish: false
salman.khan | Score: 0.179822 | Gibberish: false
michelle.smith | Score: 1.089178 | Gibberish: false
mike.johnson | Score: 1.877661 | Gibberish: false
naveen.kumar | Score: 0.139260 | Gibberish: false
michael.jones | Score: 0.813835 | Gibberish: false
suresh.kumar | Score: 0.187102 | Gibberish: false
sg4ihnio5hbqp50j9 | Score: -0.950232 | Gibberish: true
srglin5ogin5g | Score: -0.683465 | Gibberish: true
sgn4iog | Score: -1.060223 | Gibberish: true
rspghp40jg4po | Score: -1.021369 | Gibberish: true
sp4g9h4;ghs;f9s | Score: -0.949352 | Gibberish: true
srihg94hg | Score: -0.977821 | Gibberish: true
selgih4n94jgy | Score: -1.015313 | Gibberish: true
srg.ewgog | Score: -1.290902 | Gibberish: true
fxibtbq.dwf0g | Score: -0.885331 | Gibberish: true
sgl4hg9eh3h | Score: -0.794027 | Gibberish: true
vlve49gh5qh3 | Score: -0.762525 | Gibberish: true
4wfsgdrb840gu | Score: -0.794027 | Gibberish: true
fwotw45g5ugp | Score: -1.136438 | Gibberish: true
vdpt9br535 | Score: -0.794027 | Gibberish: true
ev5ot54eov | Score: -0.794027 | Gibberish: true
ldhtnhudgsvhil | Score: -1.084786 | Gibberish: true
```

## Explanation

The markov chain trains on a big set of geniune usernames (`usernames.txt`) and records the character transition probabilities.
We then calculate the transition probability scores of a set of genuine looking and gibberish usernames (`train.txt`) and the extract the mean and standard deviation of the scores. We do this to calculate the z-score of the usernames in the test set (`test.txt`)
so that they can be classified. A negative z-score indicates a gibberish username and positive one indicates a genuine one.
 