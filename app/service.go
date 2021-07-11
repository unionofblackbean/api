package app

type Service interface {
	Start() error
	Shutdown() error
	Name() string
}
