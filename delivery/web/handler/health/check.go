package health

import (
	"fmt"
	"net/http"
)

//go:generate mockery --name Check
type Check interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type check struct {
}

func NewCheck() Check {
	return &check{}
}

func (h *check) Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "ok")
}
