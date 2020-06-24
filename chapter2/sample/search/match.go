package search

import (
	"fmt"
	"log"
)

//results contains the result of a search
type Result struct {
	Field   string
	Content string
}

//matcher defines the behavior required by types that want
//to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//match is launched as a goroutines for each individual feed to run
//searches concurrently
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//perform the search against the specified matcher
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//write teh results to the channel
	for _, result := range searchResult {
		results <- result
	}
}

//display writes result to the terminal window as they
//are received byt the individual goroutine
func Display(results chan *Result) {
	//the channel blocks until a result is written to the channel
	//once the channel is closed the for loop terminates
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
