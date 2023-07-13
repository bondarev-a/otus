package main

import (
	"encoding/json"
	"net/http"
)

type checkMsg struct {
	STATUS string `json:"status"`
}

/*
func main() {
	fmt.Println("Starting...")
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })
	//router.GET("/live", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })
	//router.GET("/ready", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })
	//router.GET("/ready", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, checkMsg{STATUS: "OK"}) })

	router.Run("0.0.0.0:8000")
}*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkMsg{STATUS: "OK"})
}

func main() {
	port := "8000"
	mux := http.NewServeMux()

	mux.HandleFunc("/health", indexHandler)
	http.ListenAndServe("0.0.0.0:"+port, mux)
}
