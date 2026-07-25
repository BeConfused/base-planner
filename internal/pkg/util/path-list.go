// Package util exposes a set of utility functions to use in other packages
package util

// PathList groups all paths needed by the CLI
// Primarily used to load any specified paths required to run.
type PathList struct {
	Config string
	Plan   string
}
