package main

import (
	"log"
	"os"
	"samcrider/sql_querier/sql_querier"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// get args- should only be the json file to use as the db
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must include a JSON file as a CLI argument")
	}

	// set file path to a variable
	path := args[0]

	// get user's db query
	res := ""
	label := "SQL Query:"
	prompt := &survey.Input{
		Message: label,
	}
	survey.AskOne(prompt, &res)

	// call json parser and pass in the query and the json file path
	sql_querier.Query(path, res)
}
