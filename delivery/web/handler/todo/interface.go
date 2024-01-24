package todo

import http "net/http"

//go:generate mockery --name Handler --structname MockHandler --output . --outpkg todo --case underscore --filename handler_mock.go
type Handler interface {
	CreateTodoGroup(w http.ResponseWriter, r *http.Request)
	GetTodoGroup(w http.ResponseWriter, r *http.Request)
	GetTodoGroups(w http.ResponseWriter, r *http.Request)
	UpdateTodoGroup(w http.ResponseWriter, r *http.Request)
	DeleteTodoGroup(w http.ResponseWriter, r *http.Request)
	SortItems(w http.ResponseWriter, r *http.Request)

	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	DeleteCompletedItems(w http.ResponseWriter, r *http.Request)
}
