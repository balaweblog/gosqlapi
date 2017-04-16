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
	//No records found
	NoRecordsFound = errors.New("No Records found")
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(val) < 0 {
			warningmessage(w, NoRecordsFound.Error(), http.StatusNoContent)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(val) < 0 {
			warningmessage(w, NoRecordsFound.Error(), http.StatusNoContent)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(val) < 0 {
			warningmessage(w, NoRecordsFound.Error(), http.StatusNoContent)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagequery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
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
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err != nil {
			errormessage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		successmessagenonquery(w, val, http.StatusOK)
	})
}

/*Response json object */
type Responsequery struct {
	Status string
	Data   map[int]map[string]string
}

/*Response json object */
type Responsenonquery struct {
	Status string
	Data   interface{}
}
type ErrorResponse struct {
	Code    int
	Message string
}
type Request struct {
	Data map[string]interface{} `json:"data"`
}

func parserequest(req []byte) (Request, error) {
	var request Request
	err := json.Unmarshal(req, &request)
	if err != nil {
		return request, err
	}
	return request, nil
}

/* success response message */
func successmessagequery(w http.ResponseWriter, msg map[int]map[string]string, errorcode int) {
	response := Responsequery{Status: "success", Data: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/* success response message */
func successmessagenonquery(w http.ResponseWriter, msg interface{}, errorcode int) {
	response := Responsenonquery{Status: "success", Data: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/* error response message */
func errormessage(w http.ResponseWriter, msg string, errorcode int) {
	response := ErrorResponse{Code: errorcode, Message: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}

/* warning response message */
func warningmessage(w http.ResponseWriter, msg string, errorcode int) {
	response := ErrorResponse{Code: errorcode, Message: msg}
	message, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorcode)
	w.Write(message)
}
