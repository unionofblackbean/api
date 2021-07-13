package config

const (
	StartPolicyExitOnError StartPolicy = "exit_on_error"
	StartPolicyNeverExit   StartPolicy = "never_exit"
)

type StartPolicy string
