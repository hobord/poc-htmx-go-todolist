package todo

import http "net/http"

//go:generate mockery --name Handler --structname MockHandler --output . --outpkg todo --case underscore --filename handler_mock.go
type Handler interface {
	AddItem(w http.ResponseWriter, r *http.Request)
	TodoLists(w http.ResponseWriter, r *http.Request)
	TodoItems(w http.ResponseWriter, r *http.Request)
	TodoItem(w http.ResponseWriter, r *http.Request)
}
