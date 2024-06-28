package sql_querier

import (
	"samcrider/sql_querier/utils"
)

func Query(path string, query string) {
	selects, table, conditions := utils.Query_parser(query)

	// get the table from the database

	// if there are conditions, parse them into individual conditions
	// search through the table and grab all objects that align with the conditions
	// loop through selects
	// if selects is *, return the full objects that meet the conditions
	// else for each select, set it to values from the objects that meet the conditions

	// if no conditions
	// loop through selects
	// if selects is *, return the entire table
	// else for each select, set it to values from the objects in the table
}
