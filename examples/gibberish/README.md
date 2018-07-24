# Gibberish username detector

## Overview
`gibberish.go` is a simple program to differentiate between genuine looking and gibberish usernames. It uses a 2 character markov chain.

## Usage
```
> go run gibberish.go
Output:
david.williams - Gibberish: false
chris.williams - Gibberish: false
john.williams - Gibberish: false
andrew.smith - Gibberish: false
adam.smith - Gibberish: false
joe.smith - Gibberish: false
salman.khan - Gibberish: false
michelle.smith - Gibberish: false
mike.johnson - Gibberish: false
naveen.kumar - Gibberish: false
michael.jones - Gibberish: false
suresh.kumar - Gibberish: false
sg4ihnio5hbqp50j9 - Gibberish: true
srglin5ogin5g - Gibberish: true
sgn4iog - Gibberish: true
rspghp40jg4po - Gibberish: true
sp4g9h4;ghs;f9s - Gibberish: true
srihg94hg - Gibberish: true
selgih4n94jgy - Gibberish: true
srg.ewgog - Gibberish: true
fxibtbq.dwf0g - Gibberish: true
sgl4hg9eh3h - Gibberish: true
vlve49gh5qh3 - Gibberish: true
4wfsgdrb840gu - Gibberish: true
fwotw45g5ugp - Gibberish: true
vdpt9br535 - Gibberish: true
ev5ot54eov - Gibberish: true
ldhtnhudgsvhil - Gibberish: true
```

## Explanation

The markov chain trains on a big set of geniune usernames (`usernames.txt`) and records the character transition probabilities.
We then calculate the transition probability scores of a set of genuine looking and gibberish usernames (`train.txt`) and the extract the mean and standard deviation of the scores. We do this to calculate the z-score of the usernames in the test set (`test.txt`)
so that they can be classified. A negative z-score indicates a gibberish username and positive one indicates a genuine one.
 