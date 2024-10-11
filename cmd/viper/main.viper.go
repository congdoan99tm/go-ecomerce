package main

import (
	"fmt"
	. "github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")
	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Security Key::", viper.GetString("security.jwt.key"))
	// configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
	fmt.Println("Config Port::", viper.GetInt("config.server.port"))
	for _, db := range config.Databases {
		fmt.Printf("database User: %s, Password: %s, Host: %s\n", db.User, db.Password, db.Host)
	}
}
