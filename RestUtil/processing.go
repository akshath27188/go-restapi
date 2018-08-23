package RestUtil

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/modelproject/Model"
)

// FormToUser -- fills a User struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// User struct if successful
// array of strings of errors if any occur during processing
func FormToUser(r *http.Request) (Model.User, []string) {
	var user Model.User
	var errStr, ageStr string
	var errs []string
	var err error

	user.FirstName, errStr = ProcessFormField(r, "firstname")
	errs = AppendError(errs, errStr)
	user.LastName, errStr = ProcessFormField(r, "lastname")
	errs = AppendError(errs, errStr)
	user.Email, errStr = ProcessFormField(r, "email")
	errs = AppendError(errs, errStr)
	user.City, errStr = ProcessFormField(r, "city")
	errs = AppendError(errs, errStr)

	ageStr, errStr = ProcessFormField(r, "age")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		user.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			errs = append(errs, "Parameter 'age' not an integer")
		}
	}
	return user, errs
}

func AppendError(errs []string, errStr string) ([]string) {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func ProcessFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}

func ProcessJsonPostData(r *http.Request,q Model.User)(Model.User,[]string){
	//var user Model.User 
	var errs []string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	    errs = AppendError(errs, "Couldnt read request body")
	}
	fmt.Println("Json request data - ",string(body))
	err = json.Unmarshal(body, &q)
	//err := decoder.Decode(&user)
	if err != nil {
	     fmt.Println("Invalid Json request data",err.Error())
	     errs = AppendError(errs, "Invalid Json request data")
	}
	//fmt.Println("user Email - ", user.Email)
	return q,errs
}


