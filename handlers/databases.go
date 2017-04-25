package handlers

import (
	"gosqlapi/model"
	"net/http"
)

/*CreateddbPost - create db */
func CreatedbPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "CREATEDB")
		writenonqueryresponse(w, query)
	}
}

/*CurrentdbPost - get the current db in use */
func CurrentdbPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "CURRENTDB")
		writequeryresponse(w, query)
	}
}

/*DropdatabasePost - drop database */
func DropdatabasePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "DROPDB")
		writenonqueryresponse(w, query)
	}
}

/*ShowalldatabasesPost - show all databases */
func ShowalldatabasesPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "SHOWDB")
		writequeryresponse(w, query)
	}
}

/*showalltablesPost - show all tables */
func ShowalltablesPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "SHOWALLTABLES")
		writequeryresponse(w, query)
	}
}

type Databases struct {
}
