// Package micropub implements the micropub standard for creating, updating and
// deleting posts.
package micropub

import (
	"net/http"
)

// Service provides the dependencies required by the micropub service.
//
//encore:service
type Service struct{}

func initService() (*Service, error) {
	return &Service{}, nil
}

// Handle is the entrypoint for all micropub requests.
//
//encore:api public raw path=/micropub
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) {
	unimplemented(w, r)
}
