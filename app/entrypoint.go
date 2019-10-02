package app

import (
	"regulus/app/infrastructures/echo"
	"regulus/app/infrastructures/viper"
)

// StartApp entoru point of application
func StartApp() {
	// configuration
	viper.SetupAppConfig()
	echo.StartEcho()
}
