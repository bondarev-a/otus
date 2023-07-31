package main

import (
	"os"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/lib/pq"
)
/*
    host     = os.Getenv("db_host")
    port     = os.Getenv("db_port")
    pg_user  = os.Getenv("db_user")
    password = os.Getenv("db_password")
    dbname   = os.Getenv("db_name")
*/

type checkMsg struct {
	STATUS string `json:"status"`
}

type user struct {
    ID			int64  `json:"id"`
    UserName	string  `json:"username"`
	FirstName	string  `json:"firstName"`
	LastName	string  `json:"lastName"`
	EMail		string  `json:"email"`
	Phone		string  `json:"phone"`
}

var db *sql.DB

func initConn() *sql.DB {
    host     := os.Getenv("db_host")
    port, err:= strconv.Atoi(os.Getenv("db_port"))
    pg_user  := os.Getenv("db_user")
    password := os.Getenv("db_password")
    dbname   := os.Getenv("db_name")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, pg_user, password, dbname)

	fmt.Print(psqlconn)

	dbc, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	
	return dbc
}

func postUser(c *gin.Context){
	var userUpd user 
	jsonData, err := c.GetRawData()
	CheckError(err)
	
	fmt.Print(jsonData)

	err = json.Unmarshal(jsonData, &userUpd)
	CheckError(err)

	_, e := db.Exec("insert into users (user_name, first_name, last_name, email, phone) values($1, $2, $3, $4, $5);", userUpd.UserName, userUpd.FirstName, userUpd.LastName, userUpd.EMail, userUpd.Phone)
	
	CheckError(e)

	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func getUser(c *gin.Context){

	userRow := user{}

	err := db.QueryRow(fmt.Sprintf("SELECT user_name, first_name, last_name, email, phone FROM public.users WHERE id=%s", c.Param("userId"))).
		Scan(&userRow.UserName, &userRow.FirstName, &userRow.LastName, &userRow.EMail, &userRow.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
            c.IndentedJSON(http.StatusOK, "Error: user not found")
		}
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, userRow)
}

func deleteUser(c *gin.Context){

	delStr := fmt.Sprintf("DELETE FROM public.users WHERE id=%s", c.Param("userId"))
	_, e := db.Exec(delStr)
	
	CheckError(e)

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("user %s deleted", c.Param("userId")))
}

func putUser(c *gin.Context){
	var userUpd user 
	jsonData, err := c.GetRawData()
	CheckError(err)
	
	err = json.Unmarshal(jsonData, &userUpd)
	CheckError(err)

	insStr := fmt.Sprintf("insert into users (user_name, first_name, last_name, email, phone) values(%s, %s, %s, %s, %s);", 
						userUpd.UserName, userUpd.FirstName, userUpd.LastName, userUpd.EMail, userUpd.Phone)

	_, e := db.Exec(insStr)
	
	CheckError(e)

	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	db = initConn()

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })

	router.POST("/user", postUser)
	router.GET("/user/:userId", getUser)
	router.DELETE("/user/:userId", deleteUser)
	router.PUT("/user/:userId", putUser)

	router.Run("0.0.0.0:8000")

}
