// Package micropub implements the micropub standard for creating, updating and
// deleting posts.
package micropub

import (
	"encoding/json"
	"net/http"
	"net/url"

	"encore.dev/rlog"
)

// Service provides the dependencies required by the micropub service.
//
//encore:service
type Service struct {
	FrontendBaseURL *url.URL
	Entries         []Entry
}

func initService() (*Service, error) {
	baseURL, err := url.Parse("https://blog.example.com")
	if err != nil {
		return nil, err
	}

	return &Service{
		FrontendBaseURL: baseURL,
	}, nil
}

// Handle is the entrypoint for all micropub requests.
//
//encore:api public raw path=/micropub
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rlog.Debug("unsupported method")
		unimplemented.ServeHTTP(w, r)
		return
	}

	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		s.handleForm(w, r)
	case "application/json":
		s.handleJSON(w, r)
	}
}

func (s *Service) handleForm(w http.ResponseWriter, r *http.Request) {
	if h := r.FormValue("h"); h != "entry" {
		rlog.Debug("unsupported h type", "h", h)
		unimplemented.ServeHTTP(w, r)
		return
	}

	if err := r.ParseForm(); err != nil {
		rlog.Error("malformed body", "error", err)
		invalidRequest.ServeHTTP(w, r)
		return
	}

	// TODO: Save the entry.
	e := entryFromFormValues(r.PostForm)
	s.Entries = append(s.Entries, e)

	w.Header().Add("location", s.FrontendBaseURL.JoinPath("/entry/123").String())
	w.WriteHeader(http.StatusCreated)
}

type createRequest struct {
	Type       []string `json:"type"`
	Properties struct {
		Content []string `json:"content"`
	} `json:"properties"`
}

func (s *Service) handleJSON(w http.ResponseWriter, r *http.Request) {
	cr := &createRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(cr); err != nil {
		rlog.Error("malformed body", "error", err)
		invalidRequest.ServeHTTP(w, r)
		return
	}

	if h := cr.Type[0]; h != "h-entry" {
		rlog.Debug("unsupported h type", "h", h)
		unimplemented.ServeHTTP(w, r)
		return
	}

	// TODO: Save the entry.
	e := entryFromJSONValues(cr)
	s.Entries = append(s.Entries, e)

	w.Header().Add("location", s.FrontendBaseURL.JoinPath("/entry/123").String())
	w.WriteHeader(http.StatusCreated)
}
