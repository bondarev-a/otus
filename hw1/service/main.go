package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkMsg struct {
	STATUS string `json:"status"`
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"})
}

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })
	router.Run("localhost:8000")

}
