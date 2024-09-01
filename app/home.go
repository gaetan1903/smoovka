package app

import (
	"fmt"
	"smoovka/services"
)

// DASHBOARD
func (a *App) UserGetAccount() map[string]interface{} {
	return services.UserGetAccount()
}

func (a *App) UserGetStatus() string {
	return services.GetCurrentStatus()
}

func (a *App) UserStartSession() bool {
	return services.UserStartSession()
}

func (a *App) UserStopSession() bool {
	return services.UserStopSession()
}

func (a *App) UserCurrentTotalTime() float64 {
	times, err := services.UserTotalTimeToDay()
	if err != nil {
		fmt.Println(err)
		return 0.0
	}
	return times
}
