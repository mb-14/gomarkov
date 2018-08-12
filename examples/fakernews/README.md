# Fake HN post generator

`fakernews` builds a markov chain using the top 500 post titles on HN and generates fake HN posts

## Usage
```
> go build -o fakernews .
#Train on HN top stories
> ./fakernews -train
> ./fakernews
Output: Show HN: What Wal-Mart Knows About Customers' Habits (2004)
> ./fakernews
Output: A modern computer model could rival the Large Hadron Collider
```


