package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func query(db *sql.DB) {

	start := time.Now()
	rows, _ := db.Query("SELECT * FROM test")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d,name: %s\n", id, name)
	}
	end := time.Now()
	fmt.Println("query total time:", end.Sub(start).Seconds())

}
