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
	host := viper.GetString("yapi.db.host")
	port := viper.GetString("yapi.db.port")
	user := viper.GetString("yapi.db.user")
	password := viper.GetString("yapi.db.password")
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
	// db := viper.GetString("yapi.db")
	return viper.GetString("yapi.db.name")
}
