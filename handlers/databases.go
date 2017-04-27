package handlers

import (
	"github.com/balaweblog/gosqlapi/model"
	"net/http"
)

/*CreatedbPost - create db */
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

/*ShowalltablesPost - show all tables */
func ShowalltablesPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "SHOWALLTABLES")
		writequeryresponse(w, query)
	}
}

/*Databases - databases struct */
type Databases struct {
}
