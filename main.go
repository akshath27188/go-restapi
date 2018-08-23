package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/modelproject/DB/Cassandra"
  "github.com/modelproject/RestHandler/Users"
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
  log.Fatal(http.ListenAndServe(":8080", router))
}


func heartbeat(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}



