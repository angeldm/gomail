package gomail

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
	To          string
}

func readConfig() *Configuration {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(configuration.Users) // output: [UserA, UserB]
	return &configuration
}

func readConfigWithPath(path string) *Configuration {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(configuration.Users) // output: [UserA, UserB]
	return &configuration
}
