package main

import (
	"gosqlapi/handlers"
	"gosqlapi/model"
	"log"
	"net/http"
)

func main() {
	db, err := model.NewConnection()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	//Execute Query
	mux.Handle("/showalldatabases", handlers.Showalldatabases(db))
	mux.Handle("/currentdb", handlers.CurrentInUseDB(db))
	mux.Handle("/showalltables", handlers.ShowalltablesinDb(db))
	mux.Handle("/altertable", handlers.Altertable(db))
	mux.Handle("/describetable", handlers.Describetableindb(db))
	mux.Handle("/explaintable", handlers.ExplaintableinDb(db))
	mux.Handle("/selecttable", handlers.Selecttable(db))

	//execute query user specific queries
	mux.Handle("/listallusers", handlers.Listallusers(db))

	//execute non query with no return
	mux.Handle("/createdb", handlers.Createdb(db))
	mux.Handle("/dropdatabase", handlers.Dropdatabase(db))
	mux.Handle("/createtable", handlers.Createtable(db))
	mux.Handle("/inserttable", handlers.Inserttable(db))
	mux.Handle("/updatetable", handlers.Updatetable(db))
	mux.Handle("/deletetable", handlers.Deletetable(db))
	mux.Handle("/droptable", handlers.Droptable(db))
	mux.Handle("/truncatetable", handlers.Truncatetable(db))
	mux.Handle("/showalltableindex", handlers.Showalltableindex(db))

	log.Printf("serving on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
