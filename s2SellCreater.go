package main

import (
	"log"
	"net/http"

	"github.com/danakum/s2-cell-creator/handler"
)

func main() {

	http.HandleFunc("/createpolygon", handler.GetPolygoneAndCreateArea)
	http.HandleFunc("/getselldetails", handler.GetSellDetails)
	// Start the server at http://localhost:9988
	log.Fatal(http.ListenAndServe(":9988", nil))
}
