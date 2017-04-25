package handlers

import (
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
func writequeryresponse(w http.ResponseWriter, query string) {
	db, err := model.NewConnection()
	defer db.Close()
	if err != nil {
		errormessage(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
func writenonqueryresponse(w http.ResponseWriter,  query string) {
	db, err := model.NewConnection()
	defer db.Close()
	if err != nil {
		errormessage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val, err := model.ExecuteNonQuery(db, query)

	if err != nil {
		errormessage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successmessagenonquery(w, val, http.StatusOK)
}

/*readrequest read */
func readrequest(w http.ResponseWriter, r *http.Request) (Request, error) {
	req, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var request Request

	if len(req) < 0 || err != nil {
		errormessage(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
		return request, err
	}

	request, err = parserequest(req)

	if err != nil {
		errormessage(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
		return request, err
	}
	return request, err
}
