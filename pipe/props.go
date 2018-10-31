package pipe

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Configuration struct {
	SrcPort     int
	DestIp      string
	DestPort    int
	SerialPortA string
	SerialPortB string
}

var once sync.Once
var config *Configuration

func GetProps() *Configuration {

	once.Do(func() {
		fmt.Println("Getting properties...")
		config = &Configuration{}
		file, err := os.Open(getWd() + "/pipe/properties.json")
		if err != nil {
			panic(err)
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			panic(err)
		}
	})
	return config
}

func getWd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}
