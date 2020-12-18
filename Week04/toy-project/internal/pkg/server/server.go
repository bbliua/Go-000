package server

import "context"
// Server server interface
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
