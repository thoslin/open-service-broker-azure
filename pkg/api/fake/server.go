package fake

import (
	"context"
)

// Server is a fake implementation of api.Server used for testing
type Server struct {
	RunBehavior func() error
}

// NewServer returns a new, fake implementation of api.Server used for testing
func NewServer() *Server {
	return &Server{
		RunBehavior: defaultRunBehavior,
	}
}

// Start causes the api server to start serving HTTP requests. It will block
// until an error occurs and will return that error.
func (s *Server) Start(context.Context) error {
	return s.RunBehavior()
}

func defaultRunBehavior() error {
	select {}
}
