package goimserver

import (
    "fmt"
    "github.com/pelletier/go-toml"
)

func (c *config) loadConfig(string file) {
    config, err := toml.LoadFile("config.toml")
    if err != nil {
        fmt.Println("Error ", err.Error())
    }
}

func (c *config) getConfig(){
    
}