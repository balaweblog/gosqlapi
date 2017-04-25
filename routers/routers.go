package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"gosqlapi/handlers"
	"gosqlapi/logger"
	"net/http"
)

/*NewRouter  Routers new router */
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

/* Index  index get method */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreatedbPost",
		"POST",
		"/v1/createdb",
		handlers.CreatedbPost,
	},

	Route{
		"CurrentdbPost",
		"POST",
		"/v1/currentdb",
		handlers.CurrentdbPost,
	},

	Route{
		"DropdatabasePost",
		"POST",
		"/v1/dropdatabase",
		handlers.DropdatabasePost,
	},

	Route{
		"ShowalldatabasesPost",
		"POST",
		"/v1/showalldatabases",
		handlers.ShowalldatabasesPost,
	},

	Route{
		"ShowalltablesPost",
		"POST",
		"/v1/showalltables",
		handlers.ShowalltablesPost,
	},

	Route{
		"AltertablePost",
		"POST",
		"/v1/altertable",
		handlers.AltertablePost,
	},

	Route{
		"CreatetablePost",
		"POST",
		"/v1/createtable",
		handlers.CreatetablePost,
	},

	Route{
		"DeletetablePost",
		"POST",
		"/v1/deletetable",
		handlers.DeletetablePost,
	},

	Route{
		"DescribetablePost",
		"POST",
		"/v1/describetable",
		handlers.DescribetablePost,
	},

	Route{
		"DroptablePost",
		"POST",
		"/v1/droptable",
		handlers.DroptablePost,
	},

	Route{
		"ExplaintablePost",
		"POST",
		"/v1/explaintable",
		handlers.ExplaintablePost,
	},

	Route{
		"InserttablePost",
		"POST",
		"/v1/inserttable",
		handlers.InserttablePost,
	},

	Route{
		"ListallusersPost",
		"POST",
		"/v1/listallusers",
		handlers.ListallusersPost,
	},

	Route{
		"SelecttablePost",
		"POST",
		"/v1/selecttable",
		handlers.SelecttablePost,
	},

	Route{
		"ShowalltableindexPost",
		"POST",
		"/v1/showalltableindex",
		handlers.ShowalltableindexPost,
	},

	Route{
		"TruncatetablePost",
		"POST",
		"/v1/truncatetable",
		handlers.TruncatetablePost,
	},

	Route{
		"UpdatetablePost",
		"POST",
		"/v1/updatetable",
		handlers.UpdatetablePost,
	},
}
