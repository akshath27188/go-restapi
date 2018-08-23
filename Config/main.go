package Config

import (
  "github.com/spf13/viper"
  "fmt"
)

var Runtime_viper *viper.Viper

func init() {
  Runtime_viper = viper.New()
  Runtime_viper.SetConfigType("json")
  Runtime_viper.AddConfigPath("/home/akshath/go/src/modelproject")
  Runtime_viper.SetConfigName("config")
  Runtime_viper.ReadInConfig()
  fmt.Println("viper init done")
}

