# Pokemon name generator

`pokenamer` builds a markov chain using the existing pokemon names and generates made-up pokemin names

## Usage
```
> go build -o pokenamer .
#Train on HN top stories
> ./pokenamer -train -order 3
> ./pokenamer
Output: Cacneas
> ./pokenamer
Output: Relix
> ./pokenamer
Output: Farfet
```


