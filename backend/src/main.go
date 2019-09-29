package main

import (
	"regulus/infrastructures/echo"
	"regulus/infrastructures/viper"
)

// start apprication
func main() {
	// configuration
	viper.SetupAppConfig()
	echo.StartEcho()
}
