package app

import "fmt"

const (
	name    = "api.ubb.asia"
	version = "v0.0.0"
)

func VersionStatement() string {
	return fmt.Sprintf("%s %s", name, version)
}
