package viper

import (
	"log"

	"github.com/spf13/viper"
)

// SetupAppConfig アプリケーションの設定ファイル読み込み
func SetupAppConfig() {
	// default setting
	viper.SetDefault("mode", "production")
	viper.SetDefault("httpPort", "8080")
	viper.SetDefault("dbUrl", "127.0.0.1")
	viper.SetDefault("dbPort", "5432")
	viper.SetDefault("dbUser", "regulus")
	viper.SetDefault("dbPassword", "regulus")
	viper.SetDefault("dbName", "reglusdb")

	// read config file
	viper.SetConfigName("config")
	viper.AddConfigPath("./resources/config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Panic(err.Error())
	}
}

// SetupTestConfig test用の設定読み込み
func SetupTestConfig() {
	// default setting
	viper.SetDefault("mode", "test")
	viper.SetDefault("httpPort", "8080")
	viper.SetDefault("dbUrl", "127.0.0.1")
	viper.SetDefault("dbPort", "5432")
	viper.SetDefault("dbUser", "demo")
	viper.SetDefault("dbName", "demodb")
	viper.SetDefault("dbPassword", "demo")
}
