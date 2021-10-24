package models

import (
	log "github.com/sirupsen/logrus"
)

// App contains the objects that are required for the
// application to function correctly
type App struct {
	Config *Config
	Logger *log.Logger
}

// ConfigureLogging sets up the logging for the application
func (app *App) ConfigureLogging() {

}
