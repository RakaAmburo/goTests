package main

import (
    "fmt"
	"os"
	"encoding/json"
	"sync"
)

type Configuration struct {
    Database           DataBase
}

type DataBase struct {
	Host string
	Port string
	User string
	PassWord string
	Url string
}

var once sync.Once
var config *Configuration

func Get(key string) string{
 return ""
}

func GetProps() *Configuration {
 
	once.Do(func() {
        fmt.Println("Getting properties...")
		config = &Configuration{}
		file, err := os.Open(getWd() + "/properties.json")
		if err != nil {  panic(err) }  
		decoder := json.NewDecoder(file) 
		err = decoder.Decode(&config) 
		if err != nil {  panic(err) }
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