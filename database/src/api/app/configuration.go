package app

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type DataBaseConfig struct {
	User     string
	PassWord string
	Port     int
	Schema   string
	Url      string
}

func GetDbPropertes(path string, dbName string) *DataBaseConfig {

	viper.SetConfigName("dbConfig")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	conf := &DataBaseConfig{}
	if err != nil {
		fmt.Println("Config not found...")
		os.Exit(1)
	} else {

		if err := viper.Sub(dbName).Unmarshal(conf); err != nil {
			fmt.Printf("couldn't read config: %s", err)
			os.Exit(1)
		}

	}

	return conf
}

// properties file example named dbConfig.toml

/*
[CORE_MLB]
user = ""
password = ""
url = ""
port = 6612
schema = ""
*/
