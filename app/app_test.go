package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newApp() *App {
	deps := &Deps{}
	return New(deps)
}

func TestNew(t *testing.T) {
	app := newApp()
	assert.NotNil(t, app)
	assert.NotNil(t, app.services)
	assert.NotNil(t, app.deps)
}

type testService struct {
}

func (s *testService) Start() error {
	return nil
}

func (s *testService) Shutdown() error {
	return nil
}

func (s *testService) Name() string {
	return "test"
}

func TestApp_RegisterService(t *testing.T) {
	app := newApp()

	service := &testService{}
	app.RegisterService(service)

	assert.ElementsMatch(t, []Service{service}, app.services)
}
