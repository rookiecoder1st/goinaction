package search

import (
	"log"
	"sync"
)

//A map of registered matchers for searching

var matchers = make(map[string]Matcher)

//Run performs the search logic
func Run(searchTerm string) {
	//retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	//create a unbuffered channel to receive match results
	results := make(chan *Result)
	//setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup
	//setup a number of goroutines we need to wait for while
	//the process the individuals feeds
	waitGroup.Add(len(feeds))
	//launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		//retrieve a mather for the search
		matcher, err := matchers[feed.Type]
		if !err {
			matcher = matchers["default"]
		}
		//launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	//launch a goroutine to monitor when all the work is done
	go func() {
		//wait for everything to be processed
		waitGroup.Wait()
		//close the channel to signal to the display
		//function that we can exit the program
		close(results)
	}()
	//start  display results as the are available and
	//return after the final result is displayed
	Display(results)
}
func Register(feedType string, matcher Matcher) {
	if _, err := matchers[feedType]; err {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
