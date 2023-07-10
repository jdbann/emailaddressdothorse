package micropub

import (
	"encoding/json"
	"net/http"
)

// Error is the response data for a request which could not be completed
// successfully.
type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// func unimplemented(w http.ResponseWriter, _ *http.Request) {
// 	w.Header().Add("content-type", "application/json")
// 	w.WriteHeader(http.StatusNotImplemented)
// 	enc := json.NewEncoder(w)
// 	_ = enc.Encode(Error{
// 		Error:            "unimplemented",
// 		ErrorDescription: "implementation cannot yet handle this request",
// 	})
// }

func errorHandler(status int, msg, desc string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		enc := json.NewEncoder(w)
		_ = enc.Encode(Error{
			Error:            msg,
			ErrorDescription: desc,
		})
	})
}

var (
	unimplemented  = errorHandler(http.StatusNotImplemented, "unimplemented", "implementation cannot yet handle this request")
	invalidRequest = errorHandler(http.StatusBadRequest, "invalid_request", "request is missing a required parameter, or there was a problem with a value of one of the parameters")
)
