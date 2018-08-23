package Users

import (
  "net/http"
  "github.com/gocql/gocql"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/modelproject/Model"
  "github.com/modelproject/DB/Cassandra"
  "fmt"
)

 // "github.com/modelproject/RestUtil"
func DeleteById(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
  var errs []string

  //vars := mux.Vars(r)
  id := ps.ByName("user_uuid")

  uuid, err := gocql.ParseUUID(id)
  if err != nil {
    errs = append(errs, err.Error())
  } else {
    query := "DELETE FROM users WHERE id=?"
    if err := Cassandra.Session.Query(query, uuid).Exec(); err !=nil {
	errs = append(errs,err.Error())
	json.NewEncoder(w).Encode(Model.ErrorResponse{Errors: errs})
    }else
    {
	fmt.Println("user_id", uuid)
	json.NewEncoder(w).Encode(Model.NewUserResponse{ID: uuid})
    }
  }
}

