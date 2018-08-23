package Users

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"modelproject/DB/Cassandra"
"modelproject/RestUtil"
"modelproject/Model"
"github.com/julienschmidt/httprouter"
"fmt"
)


func Post(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
  var errs []string
  var gocqlUuid gocql.UUID
  var user Model.User
  // FormToUser() is included in Users/processing.go
  // we will describe this later
  //user, errs := FormToUser(r)
  user,errs = RestUtil.ProcessJsonPostData(r,user)

  // have we created a user correctly
  var created bool = false

  // if we had no errors from FormToUser, we will
  // attempt to save our data to Cassandra
  if len(errs) == 0 {
    fmt.Println("creating a new user")

    // generate a unique UUID for this user
    gocqlUuid = gocql.TimeUUID()

    // write data to Cassandra
    if err := Cassandra.Session.Query(`
      INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)`,
      gocqlUuid, user.FirstName, user.LastName, user.Email, user.City, user.Age).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
    }
  }

  // depending on whether we created the user, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    fmt.Println("user_id", gocqlUuid)
    json.NewEncoder(w).Encode(Model.NewUserResponse{ID: gocqlUuid})
  } else {
    fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(Model.ErrorResponse{Errors: errs})
  }
}


func Update(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
  var errs []string
  var user Model.User
  // FormToUser() is included in Users/processing.go
  // we will describe this later
  //user, errs := FormToUser(r)
  user,errs = RestUtil.ProcessJsonPostData(r,user)
  queryValues := r.URL.Query()
  id := queryValues.Get("uuid")
  uuid, err := gocql.ParseUUID(id)

  // have we created a user correctly
  var updated bool = false

  // if we had no errors from FormToUser, we will
  // attempt to save our data to Cassandra
  if len(errs) == 0 {
    fmt.Println("update existing user")
         // generate a unique UUID for this user
     //fmt.Println("query param uuid ",uuid)
    // write data to Cassandra
    if(err != nil){
	errs = append(errs, err.Error())
    } else {
	fmt.Println("about to update user")
       if err = Cassandra.Session.Query(`
	 UPDATE users set firstname=?,lastname=?,email=?,city=?,age=? where id=?`,
	 user.FirstName, user.LastName, user.Email, user.City, user.Age,id).Exec(); err != nil {
	 errs = append(errs, err.Error())
       } else {
	 updated = true
       }
    }
  }


  // depending on whether we created the user, return the
  // resource ID in a JSON payload, or return our errors
  if updated {
    fmt.Println("user_id", uuid)
    json.NewEncoder(w).Encode(Model.NewUserResponse{ID: uuid})
  } else {
     fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(Model.ErrorResponse{Errors: errs})
  }
}

