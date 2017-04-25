package handlers

import (
	"gosqlapi/model"
	"net/http"
)

/*AltertablePost - alter table  */
func AltertablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "ALTER")
		writequeryresponse(w, query)
	}
}

/*CreatetablePost - create table  */
func CreatetablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "CREATE")
		writenonqueryresponse(w, query)
	}
}

/*DeletetablePost - delete table  */
func DeletetablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "DELETE")
		writenonqueryresponse(w, query)
	}
}

/*DescribetablePost - describe table */
func DescribetablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "DESCTABLE")
		writequeryresponse(w, query)
	}
}

/*DroptablePost - drop table */
func DroptablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "DROPTABLE")
		writenonqueryresponse(w, query)
	}
}

/*ExplaintablePost - explain table */
func ExplaintablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "EXPLAINTABLE")
		writequeryresponse(w, query)
	}
}

/*InserttablePost - insert table */
func InserttablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "INSERT")
		writenonqueryresponse(w, query)
	}
}

/*ListallusersPost - list all users */
func ListallusersPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "LISTALLUSERS")
		writequeryresponse(w, query)
	}
}

/*SelecttablePost - select table */
func SelecttablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "SELECT")
		writequeryresponse(w, query)
	}
}

/*ShowalltableindexPost - show all table index */
func ShowalltableindexPost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "SHOWALLTABLEINDEX")
		writenonqueryresponse(w, query)
	}
}

/*TruncatetablePost - truncate table */
func TruncatetablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "TRUNCATE")
		writenonqueryresponse(w, query)
	}
}

/*UpdatetablePost - update table */
func UpdatetablePost(w http.ResponseWriter, r *http.Request) {
	request, err := readrequest(w, r)
	if err != nil {
		query := model.ParseQuery(request.Data, "UPDATE")
		writenonqueryresponse(w, query)
	}
}

/*Tables - table struct */
type Tables struct {
}
