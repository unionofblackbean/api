package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/unionofblackbean/api/app/config"
	"testing"
)

func TestNew(t *testing.T) {
	app := New(&Deps{})
	assert.NotNil(t, app)
	assert.NotNil(t, app.deps)
}

type registerTestService struct {
}

func (s *registerTestService) Start() error {
	return nil
}

func (s *registerTestService) Shutdown() error {
	return nil
}

func (s *registerTestService) Name() string {
	return ""
}

func TestApp_RegisterService(t *testing.T) {
	app := New(&Deps{})

	service := new(registerTestService)
	app.RegisterService(service)

	assert.ElementsMatch(t, []Service{service}, app.services)
}

type exitOnErrorStartTestService struct {
}

func (s *exitOnErrorStartTestService) Start() error {
	return errors.New("test error for exit on error policy")
}

func (s *exitOnErrorStartTestService) Shutdown() error {
	return nil
}

func (s *exitOnErrorStartTestService) Name() string {
	return ""
}

func TestApp_Start_ExitOnError(t *testing.T) {
	app := New(&Deps{
		Config: &config.Config{
			App: &config.AppConfig{
				StartPolicy: config.StartPolicyExitOnError,
			},
		},
	})
	app.RegisterService(new(exitOnErrorStartTestService))

	err := app.Start()
	assert.NotNil(t, err)
}

type neverExitStartTestService struct {
}

func (s *neverExitStartTestService) Start() error {
	return errors.New("test error for never exit policy")
}

func (s *neverExitStartTestService) Shutdown() error {
	return nil
}

func (s *neverExitStartTestService) Name() string {
	return ""
}

func TestApp_Start_NeverExit(t *testing.T) {
	app := New(&Deps{
		Config: &config.Config{
			App: &config.AppConfig{
				StartPolicy: config.StartPolicyNeverExit,
			},
		},
	})
	app.RegisterService(new(neverExitStartTestService))

	err := app.Start()
	assert.Nil(t, err)
}

type shutdownNoErrorService struct {
}

func (s *shutdownNoErrorService) Start() error {
	return nil
}

func (s *shutdownNoErrorService) Shutdown() error {
	return nil
}

func (s *shutdownNoErrorService) Name() string {
	return ""
}

func TestApp_Shutdown_NoError(t *testing.T) {
	app := New(&Deps{})
	app.RegisterService(new(shutdownNoErrorService))

	errs := app.Shutdown()
	assert.Equal(t, 0, len(errs))
}

type shutdownWithErrorService struct {
}

func (s *shutdownWithErrorService) Start() error {
	return nil
}

func (s *shutdownWithErrorService) Shutdown() error {
	return errors.New("test error")
}

func (s *shutdownWithErrorService) Name() string {
	return ""
}

func TestApp_Shutdown_WithError(t *testing.T) {
	app := New(&Deps{})
	app.RegisterService(new(shutdownWithErrorService))

	errs := app.Shutdown()
	assert.Equal(t, 1, len(errs))
}
