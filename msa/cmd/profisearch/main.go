package main

import (
	//"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)

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

func postUser(c *gin.Context){
	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func getUser(c *gin.Context){
	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func deleteUser(c *gin.Context){
	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func putUser(c *gin.Context){
	c.IndentedJSON(http.StatusOK, user{ID: 0, UserName: "put Method"})
}

func main() {

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })

	router.POST("/user", postUser)
	router.GET("/user/:userId", getUser)
	router.DELETE("/user/:userId", deleteUser)
	router.PUT("/user/:userId", putUser)

	router.Run("0.0.0.0:8000")
}
