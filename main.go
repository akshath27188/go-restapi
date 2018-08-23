package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "modelproject/Config"
  "modelproject/DB/Cassandra"
  "modelproject/RestHandler/Users"
)
//"github.com/julienschmidt/httprouter"
type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}


func main(){
  //initialized by init func in Cassandra
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()

  //var runtime_viper = viper.New()
  //viper.SetConfigType("json")
  //viper.AddConfigPath(".")
  //viper.SetConfigName("config.json")
  //router := mux.NewRouter().StrictSlash(true)
  router := httprouter.New()
  //router.HandleFunc("/", heartbeat)
  router.HandleMethodNotAllowed=false
  router.GET("/", heartbeat)
  router.POST("/users/new", Users.Post)
  router.POST("/users/update",Users.Update)
  router.GET("/users", Users.Get)
  router.GET("/users/:user_uuid", Users.GetOne)
  router.DELETE("/users/:user_uuid", Users.DeleteById)
  log.Fatal(http.ListenAndServe(":"+Config.Runtime_viper.GetString("rest_node.port"), router))
}


func heartbeat(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}



