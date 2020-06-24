package search

//defaultMatcher implement the default matcher
type defaultMatcher struct {
}

//init register the default matcher with the program
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

//search implement the behavior for the default matcher
func (self defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
