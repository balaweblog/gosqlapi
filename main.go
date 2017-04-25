package main

import (
	"gosqlapi/routers"
	"log"
	"net/http"
)

/*main  entrant method */
func main() {
	log.Printf("serving on port 8080")
	router := routers.NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
