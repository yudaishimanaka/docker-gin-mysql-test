package main

import (
	"fmt"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id	 int
	Name string
}

func main() {
	// mysql
	engine, err := sql.Open("mysql", "root:password@tcp(db-server:3306)/test_db")
	defer engine.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// gin-gonic
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		results, err := engine.Query("SELECT * FROM test_table WHERE id = 1")
		if err != nil {
			fmt.Println(err)
		}

		user := User{}
		for results.Next() {
			err = results.Scan(&user.Id, &user.Name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(user)
		}

		c.JSON(200, gin.H{
			"ID": user.Id,
			"Name": user.Name,
		})

	})

	r.Run(":8888")
}
