package index

import http "net/http"

//go:generate mockery --name Handler --structname MockHandler --output . --outpkg index --case underscore --filename handler_mock.go
type Handler interface {
	IndexPage(w http.ResponseWriter, r *http.Request)
}
