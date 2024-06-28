package utils

import (
	"log"
	"strings"
)

func Query_parser(query string) ([]string, string, []string) {
	// split the query string
	query_array := strings.Split(query, " ")

	// init flags
	select_flag := false
	from_flag := false
	where_flag := false

	// init selects
	var selects []string
	// init from
	var table string
	// init conditions
	var conditions []string
	// search args for --depth flag
	for i := range query_array {
		// set flags
		switch query_array[i] {
		case "SELECT":
			select_flag = true
			where_flag = false
			from_flag = false
			continue
		case "FROM":
			select_flag = false
			where_flag = false
			from_flag = true
			continue
		case "WHERE":
			select_flag = false
			where_flag = true
			from_flag = false
			continue
		}

		if select_flag {
			// all values get appended to selects array
			selects = append(selects, query_array[i])
		} else if from_flag {
			// there can only be one from
			table = query_array[i]
		} else if where_flag {
			// all values get appended to conditions array
			conditions = append(conditions, query_array[i])
		} else {
			log.Fatal("Query must be formatted as SQL")
		}

	}

	return selects, table, conditions
}
