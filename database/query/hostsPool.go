package query

import (
	"database/sql"
	"fmt"
	mydb "llgin/database/init"
)

//GetAllNodes获取集群信息列表
func GetAllNodes() []string {

	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT host_name FROM hosts_pool")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	var value []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}
		for _, col := range values {
			if col == nil {
				value = append(value, "NULL")
			} else {
				value = append(value, string(col))
			}
		}
	}
	return value
}
