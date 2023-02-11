package util

import (
	"context"
	"database/sql"
	"fmt"
)

func mysqlConnection(transString string) (*sql.Rows, error) {
	fmt.Println("Get Recipes")
	db, err := sql.Open("mysql", "root:Chester89!@tcp(127.0.0.1:3306)/bunkyrecipedb")
	ctx := context.Background()

	tsql := fmt.Sprintf(transString)
	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}
