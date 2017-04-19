package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"CreatedbPost",
		"POST",
		"/v1/createdb",
		CreatedbPost,
	},

	Route{
		"CurrentdbPost",
		"POST",
		"/v1/currentdb",
		CurrentdbPost,
	},

	Route{
		"DropdatabasePost",
		"POST",
		"/v1/dropdatabase",
		DropdatabasePost,
	},

	Route{
		"ShowalldatabasesPost",
		"POST",
		"/v1/showalldatabases",
		ShowalldatabasesPost,
	},

	Route{
		"ShowalltablesPost",
		"POST",
		"/v1/showalltables",
		ShowalltablesPost,
	},

	Route{
		"AltertablePost",
		"POST",
		"/v1/altertable",
		AltertablePost,
	},

	Route{
		"CreatetablePost",
		"POST",
		"/v1/createtable",
		CreatetablePost,
	},

	Route{
		"DeletetablePost",
		"POST",
		"/v1/deletetable",
		DeletetablePost,
	},

	Route{
		"DescribetablePost",
		"POST",
		"/v1/describetable",
		DescribetablePost,
	},

	Route{
		"DroptablePost",
		"POST",
		"/v1/droptable",
		DroptablePost,
	},

	Route{
		"ExplaintablePost",
		"POST",
		"/v1/explaintable",
		ExplaintablePost,
	},

	Route{
		"InserttablePost",
		"POST",
		"/v1/inserttable",
		InserttablePost,
	},

	Route{
		"ListallusersPost",
		"POST",
		"/v1/listallusers",
		ListallusersPost,
	},

	Route{
		"SelecttablePost",
		"POST",
		"/v1/selecttable",
		SelecttablePost,
	},

	Route{
		"ShowalltableindexPost",
		"POST",
		"/v1/showalltableindex",
		ShowalltableindexPost,
	},

	Route{
		"TruncatetablePost",
		"POST",
		"/v1/truncatetable",
		TruncatetablePost,
	},

	Route{
		"UpdatetablePost",
		"POST",
		"/v1/updatetable",
		UpdatetablePost,
	},

}