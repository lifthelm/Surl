package interfaces

type HealthCheck interface {
	// Name returns name of the check
	Name() string
	// Exec executes a single health check, returns details as any (optional) and error if check fails
	Exec() (any, error)
}
