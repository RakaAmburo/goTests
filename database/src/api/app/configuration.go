package app

import (
	"github.com/mercadolibre/goTests/database/src/api/app/tools"
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

func GetDbProperties(path string, dbName string) *DataBaseConfig {

	viper.SetConfigName("dbConfig")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	conf := &DataBaseConfig{}
	if err != nil {
		tools.CheckError("Config not found...", err)
		os.Exit(1)
	} else {

		if err := viper.Sub(dbName).Unmarshal(conf); err != nil {
			tools.CheckError("couldn't read config: ", err)
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
