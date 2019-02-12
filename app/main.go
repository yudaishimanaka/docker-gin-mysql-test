package main

import (
	"fmt"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

func main() {
	// mysql
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/test-db")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// gin-gonic
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM test_table")
		if err != nil {
			fmt.Println(err)
		}

		user := User{}
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(user)
		}

		c.JSON(200, gin.H{
			"Id": user.Id,
			"Name": user.Name,
		})
	})

	r.Run(":8888")
}
