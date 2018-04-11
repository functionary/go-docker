package api // import "docker.io/go-docker/api"

// Common constants for daemon and client.
const (
	// DefaultVersion of Current REST API
	DefaultVersion string = "1.33"

	// NoBaseImageSpecifier is the symbol used by the FROM
	// command to specify that no base image is to be used.
	NoBaseImageSpecifier string = "scratch"
)
