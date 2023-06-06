package main

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkMsg struct {
	STATUS string `json:"status"`
}

func main() {
	fmt.Println("Starting...")
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })
	router.Run("0.0.0.0:8000")
}
