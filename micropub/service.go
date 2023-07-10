// Package micropub implements the micropub standard for creating, updating and
// deleting posts.
package micropub

import (
	"net/http"
	"net/url"

	"encore.dev/rlog"
)

// Service provides the dependencies required by the micropub service.
//
//encore:service
type Service struct {
	FrontendBaseURL *url.URL
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
		unimplemented(w, r)
		return
	}

	if h := r.FormValue("h"); h != "entry" {
		rlog.Debug("unsupported h type", "h", h)
		unimplemented(w, r)
		return
	}

	content := r.FormValue("content")

	// TODO: Save the entry.
	rlog.Debug("create entry", "content", content)

	w.Header().Add("location", s.FrontendBaseURL.JoinPath("/entry/123").String())
	w.WriteHeader(http.StatusCreated)
}
