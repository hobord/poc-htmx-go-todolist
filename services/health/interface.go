package health

// Service is the interface for health check
//
//go:generate mockery --name Service --structname MockService --output . --outpkg health --case underscore --filename service_mock.go
type Service interface {
	Health() error
	AddChecker(func() error)
}
