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

func unimplemented(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	enc := json.NewEncoder(w)
	_ = enc.Encode(Error{
		Error:            "unimplemented",
		ErrorDescription: "implementation cannot yet handle this request",
	})
}
