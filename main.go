package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"database/sql"
    "fmt"
    _ "time"
    _ "github.com/mattn/go-sqlite3"
)

type KeyValPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetResponse struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Timestamp string `json:"timestamp"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (getResponse *GetResponse) AddRow(item Row) []Row {
    getResponse.Rows = append(getResponse.Rows, item)
    return getResponse.Rows
}

func main() {
	db, err := sql.Open("sqlite3", "./web-api.db")
	checkErr(err)
	server := gin.Default()
	server.GET("/list", func (c *gin.Context){
		rows, err := db.Query("SELECT ts, key, val FROM my_db ORDER BY ts DESC")
		checkErr(err)
		res := GetResponse{}
		for rows.Next() {
			row := Row{}
			err = rows.Scan(&row.Timestamp, &row.Key, &row.Value)
			checkErr(err)
			res.AddRow(row)
		}
		c.JSON(200, res.Rows)

	})

	server.POST("/add", func (c *gin.Context){
		json := KeyValPair{}
		c.BindJSON(&json)
		stmt, err := db.Prepare("INSERT or REPLACE INTO my_db(key, val) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(json.Key, json.Value)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("id:", id)

		c.JSON(200, gin.H{
			"key": json.Key,
			"value": json.Value,
		})
	})

	server.Run(":80")
}