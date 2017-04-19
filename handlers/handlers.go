package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"gosqlapi/model"
	"io/ioutil"
	"net/http"
)

var (
	// ErrInvalidInputRequest Error Invalid Request
	ErrInvalidInputRequest = errors.New("http: error reading input request")
	//ErrInvalidOutputResponse Error Invalid Output Response
	ErrInvalidOutputResponse = errors.New("http: error formatting response")
	//ErrParsingInputRequest Error Parsing Input Request
	ErrParsingInputRequest = errors.New("http: error parsing input request")
	//ErrNoRecordsFound No records found
	ErrNoRecordsFound = errors.New("No Records found")
)

/*Showalldatabases Show all databases in the sql */
func Showalldatabases(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "SHOWDB")
		writequeryresponse(db, query)
	})
}

/*CurrentInUseDB what database in use */
func CurrentInUseDB(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "CURRENTDB")
		val, err := model.ExecuteQuery(db, query)
		writequeryresponse(db, query)
	})
}

/*ShowalltablesinDb show all tables in the database */
func ShowalltablesinDb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "SHOWALLTABLES")
		writequeryresponse(db, query)
	})
}

/*Altertable alter table and columns in  sql */
func Altertable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "ALTER")
		writequeryresponse(db, query)
	})
}

/*ExplaintableinDb explain table in sql */
func ExplaintableinDb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "EXPLAINTABLE")
		writequeryresponse(db, query)
	})
}

/*Describetableindb describle individual table in the database */
func Describetableindb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "DESCTABLE")
		writequeryresponse(db, query)
	})
}

/*Selecttable select from table in sql */
func Selecttable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "SELECT")
		writequeryresponse(db, query)
	})
}

/*Listallusers select from table in sql */
func Listallusers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "LISTALLUSERS")
		writequeryresponse(db, query)
	})
}

/*Createdb create db in sql */
func Createdb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}
		query := model.ParseQuery(request.Data, "CREATEDB")
		writenonqueryresponse(db, query)

	})
}

/*Dropdatabase - drop database */
func Dropdatabase(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "DROPDB")
		writenonqueryresponse(db, query)
	})
}

/*Createtable create table in sql */
func Createtable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "CREATE")
		writenonqueryresponse(db, query)
	})
}

/*Inserttable insert into table in sql */
func Inserttable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "INSERT")
		writenonqueryresponse(db, query)
	})
}

/*Updatetable update into table in sql */
func Updatetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "UPDATE")
		writenonqueryresponse(db, query)
	})
}

/*Deletetable delete into table in sql */
func Deletetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "DELETE")
		writenonqueryresponse(db, query)
	})
}

/*Droptable - drop table */
func Droptable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "DROPTABLE")
		writenonqueryresponse(db, query)
	})
}

/*Truncatetable  db in sql */
func Truncatetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "TRUNCATE")
		writenonqueryresponse(db, query)
	})
}

/*Showalltableindex create db in sql */
func Showalltableindex(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if len(req) < 0 || err != nil {
			errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		request, err := parserequest(req)

		if err != nil {
			errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(request.Data, "SHOWALLTABLEINDEX")
		writenonqueryresponse(db, query)
	})
}

/*Responsequery Response json object */
type Responsequery struct {
	Status string
	Data   map[int]map[string]string
}

/*Responsenonquery Response json object */
type Responsenonquery struct {
	Status string
	Data   interface{}
}

/*ErrorResponse error repsonse */
type ErrorResponse struct {
	Code    int
	Message string
}

/*Request  input request json*/
type Request struct {
	Data map[string]interface{} `json:"data"`
}

/*parserequest -parse input request*/
func parserequest(req []byte) (Request, error) {
	var request Request
	err := json.Unmarshal(req, &request)
	if err != nil {
		return request, err
	}
	return request, nil
}

/*successmessagequery success response message */
func successmessagequery(w http.ResponseWriter, msg map[int]map[string]string, errorcode int) {
	response := Responsequery{Status: "success", Data: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/*successmessagenonquery success response message */
func successmessagenonquery(w http.ResponseWriter, msg interface{}, errorcode int) {
	response := Responsenonquery{Status: "success", Data: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/*errormessage error response message */
func errormessage(w http.ResponseWriter, msg string, errorcode int) {
	response := ErrorResponse{Code: errorcode, Message: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/*warningmessage warning response message */
func warningmessage(w http.ResponseWriter, msg string, errorcode int) {
	response := ErrorResponse{Code: errorcode, Message: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/*writequeryresponse write response */
func writequeryresponse(db *sql.DB, query string) {
	val, err := model.ExecuteQuery(db, query)

	if err != nil {
		errormessage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(val) < 0 {
		warningmessage(w, ErrNoRecordsFound.Error(), http.StatusNoContent)
		return
	}
	successmessagequery(w, val, http.StatusOK)
}

/* writenonqueryresponse non query response */
func writenonqueryresponse(db *sql.DB, query string) {
	val, err := model.ExecuteNonQuery(db, query)

	if err != nil {
		errormessage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successmessagenonquery(w, val, http.StatusOK)
}
