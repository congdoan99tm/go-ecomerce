package initialize

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	. "github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
