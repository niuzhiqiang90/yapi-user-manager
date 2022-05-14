/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetMongoUri() string {
	InitConfig()
	host := viper.GetString("apidoc.host")
	port := viper.GetString("apidoc.port")
	user := viper.GetString("apidoc.user")
	password := viper.GetString("apidoc.password")
	mongoUri := ""
	if user == "" || password == "" {
		mongoUri = "mongodb://" + host + ":" + port

	} else {
		mongoUri = "mongodb://" + user + ":" + password + "@" + host + ":" + port

	}
	return mongoUri
}

func GetDBName() string {
	InitConfig()
	// db := viper.GetString("apidoc.db")
	return viper.GetString("apidoc.db")
}
