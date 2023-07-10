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
		unimplemented.ServeHTTP(w, r)
		return
	}

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
	rlog.Debug("create entry", "entry", entryFromFormValues(r.PostForm))

	w.Header().Add("location", s.FrontendBaseURL.JoinPath("/entry/123").String())
	w.WriteHeader(http.StatusCreated)
}

type entry struct {
	Content    string
	Categories []string
	Photo      string
}

func entryFromFormValues(form url.Values) *entry {
	return &entry{
		Content:    form.Get("content"),
		Categories: form["category[]"],
		Photo:      form.Get("photo"),
	}
}
