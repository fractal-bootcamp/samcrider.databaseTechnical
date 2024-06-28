package sql_querier

import (
	"encoding/json"
	"log"
	"os"
	"samcrider/sql_querier/utils"
)

func load_JSON[T any](filename string) (T, error) {
	var data T
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func Query(path string, query string) {
	selects, table, _ := utils.Query_parser(query)

	// get the json from the path
	database, err := load_JSON[map[string]interface{}](path)
	if err != nil {
		log.Fatal("JSON couldn't be loaded from path:", path)
	}

	// get the table from the database
	table_from_db := database[table]
	asserted_table := table_from_db.([]interface{})

	// if there are conditions, parse them into individual conditions
	// search through the table and grab all objects that align with the conditions
	// loop through selects
	// if selects is *, return the full objects that meet the conditions
	// else for each select, set it to values from the objects that meet the conditions

	// if no conditions
	// loop through selects
	for i := range selects {
		// if selects is *, return the entire table
		if selects[i] == "*" {
			log.Println(table_from_db)
		} else {
			// else for each select, set it to values from the objects in the table
			for j := range asserted_table {
				asserted_row := asserted_table[j].(map[string]interface{})
				log.Println(asserted_row[selects[i]])
			}
		}
	}
}
