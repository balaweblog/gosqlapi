package main

import (
	"handlers"
	"model"
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
	mux.Handle("/showalldatabases", handlers.Showalldatabases(db))
	mux.Handle("/createtable", handlers.Createtable(db))
	mux.Handle("/selecttable", handlers.Selecttable(db))
	mux.Handle("/inserttable", handlers.Inserttable(db))
	mux.Handle("/currentdb", handlers.CurrentInUseDB(db))
	mux.Handle("/showalltables", handlers.ShowalltablesinDb(db))
	mux.Handle("/describetable", handlers.Describetableindb(db))
	mux.Handle("/explaintable", handlers.ExplaintableinDb(db))
	mux.Handle("/altertable", handlers.Altertable(db))

	log.Printf("serving on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

/*
create table request
{
	"table": "employee2422",
	"column": [{
		"first": "varchar(30)"
	}, {
		"second": "int"
	}]
}

select table request

{
	"table": "employee2422",
	"column": ["first","second"]
}


/*
{"table": "welcome",
	"column": [{
		"0": "name1",
		"1":"age1"
	}],

"properties":[{
	"name1":"varchar(20)",
	"age1":"int not null"
}]
}
*/

/*
{"table": "welcome",
	"column": [{
		"0": "name",
		"1":"age"
	}],
	"where":[
		{
			"StartTag":"(",
			"LHS":"name",
			"Operator":"=",
			"RHS":"'balamurugan'",
			"NextOp":"AND",
			"EndTag":""
		},
		{
			"StartTag":"",
			"LHS":"age",
			"Operator":">",
			"RHS":"2",
			"NextOp":"",
			"EndTag":")"
		}
		]
}
*/

/*
{"table": "welcome",
	"column": [{
		"0": "name1",
		"1":"age1"
	}],

"properties":[{
	"name1":"varchar(20)",
	"age1":"int not null"
}]
}
*/
