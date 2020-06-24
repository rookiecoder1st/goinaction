package main

import (
	_ "goinaction/chapter2/sample/matchers"
	"goinaction/chapter2/sample/search"
	"log"
	"os"
)

//init is called prior to main
func init() {
	//change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

//main is the entry point for the program
func main() {
	//perform the search for the specified  term
	search.Run("president")
}
