package database

import (
	"fmt"
	"llgin/structdb"
	"log"
	"time"
)

func Query1() {

	db := InitDB()
	defer db.Close()

	start := time.Now()

	rows, _ := db.Query("SELECT * FROM test")
	defer rows.Close()

	for rows.Next() {
		var a structdb.ClusterPool

		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d,name: %s\n", a.ID, a.Name)
	}
	end := time.Now()
	fmt.Println("query total time:", end.Sub(start).Seconds())
}
