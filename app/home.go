package app

import "smoovka/services"

// DASHBOARD
func (a *App) UserGetAccount() map[string]interface{} {
	return services.UserGetAccount()
}

func (a *App) UserGetStatus() string {
	return services.GetCurrentStatus()
}
