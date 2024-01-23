package health

// Service is the interface for health check
//
//go:generate mockery --name Service --structname MockService --inpackage --case underscore --disable-version-string
type Service interface {
	Health() error
	AddChecker(func() error)
}
