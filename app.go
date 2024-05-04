package main

import (
	"context"
	"smoovka/services"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AUTHENTICATION
func (a *App) UserIsLoggedIn() bool {
	return services.UserIsLoggedIn()
}

func (a *App) UserLogin(login string, password string) bool {
	return services.UserLogin(login, password)
}

func (a *App) UserLogout() {
	services.UserLogout()
}

// ------------------------

// DASHBOARD
func (a *App) UserGetAccount() map[string]interface{} {
	return services.UserGetAccount()
}

// ------------------------
