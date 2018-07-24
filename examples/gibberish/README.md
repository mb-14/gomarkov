# Gibberish username detector

## Overview
`gibberish.go` is a simple program to differentiate between genuine looking and gibberish usernames. It uses a 2 character markov chain.

## Usage
```
> go run gibberish.go
Output:
david.williams | Score: 0.709420 | Gibberish: false
chris.williams | Score: 0.803761 | Gibberish: false
john.williams | Score: 0.794337 | Gibberish: false
andrew.smith | Score: 0.277211 | Gibberish: false
adam.smith | Score: 0.249324 | Gibberish: false
joe.smith | Score: 0.364269 | Gibberish: false
salman.khan | Score: 0.820657 | Gibberish: false
michelle.smith | Score: 1.291605 | Gibberish: false
mike.johnson | Score: 0.831674 | Gibberish: false
naveen.kumar | Score: 0.368860 | Gibberish: false
michael.jones | Score: 1.251859 | Gibberish: false
suresh.kumar | Score: 0.323584 | Gibberish: false
sg4ihnio5hbqp50j9 | Score: -1.110799 | Gibberish: true
srglin5ogin5g | Score: -0.351409 | Gibberish: true
sgn4iog | Score: -1.274980 | Gibberish: true
rspghp40jg4po | Score: -0.900493 | Gibberish: true
sp4g9h4;ghs;f9s | Score: -0.040475 | Gibberish: true
srihg94hg | Score: -1.348694 | Gibberish: true
selgih4n94jgy | Score: -0.745763 | Gibberish: true
srg.ewgog | Score: -1.335064 | Gibberish: true
fxibtbq.dwf0g | Score: -1.322951 | Gibberish: true
sgl4hg9eh3h | Score: -1.106566 | Gibberish: true
vlve49gh5qh3 | Score: -0.051361 | Gibberish: true
4wfsgdrb840gu | Score: -1.075832 | Gibberish: true
fwotw45g5ugp | Score: -1.113248 | Gibberish: true
vdpt9br535 | Score: -1.099270 | Gibberish: true
ev5ot54eov | Score: -0.514665 | Gibberish: true
ldhtnhudgsvhil | Score: -1.319490 | Gibberish: true

```

## Explanation

The markov chain trains on a big set of geniune usernames (`usernames.txt`) and records the character transition probabilities.
We then calculate the transition probability scores of a set of genuine looking and gibberish usernames (`train.txt`) and the extract the mean and standard deviation of the scores. We do this to calculate the z-score of the usernames in the test set (`test.txt`)
so that they can be classified. A negative z-score indicates a gibberish username and positive one indicates a genuine one.
 