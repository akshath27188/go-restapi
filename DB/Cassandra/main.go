package Cassandra

import (
  "modelproject/Config"
  "github.com/gocql/gocql"
  "fmt"
)

var Session *gocql.Session

func init() {
  var err error
  //Viper.viper.SetConfigType("json")
  //Viper.viper.AddConfigPath(".")
  //Viper.viper.setConfigName("config.json")
  host := Config.Runtime_viper.GetString("db_node.host")
  fmt.Println("HOST",host)
  cluster := gocql.NewCluster(host)
  cluster.Keyspace = "user_detail"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}
