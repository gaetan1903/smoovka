package app

import (
	"smoovka/services"
)

func (a *App) UserIsLoggedIn() bool {
	return services.UserIsLoggedIn()
}

func (a *App) UserLogin(login string, password string) bool {
	return services.UserLogin(login, password)
}

func (a *App) UserLogout() {
	services.UserLogout()
}
