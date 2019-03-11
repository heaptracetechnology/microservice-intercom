package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-intercom/result"
	intercom "gopkg.in/intercom/intercom-go.v2"
	"net/http"
	"os"
)

//Create User
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	ic := intercom.NewClient(accessToken, "")

	decoder := json.NewDecoder(request.Body)
	var param *intercom.User
	err := decoder.Decode(&param)
	savedUser, err := ic.Users.Save(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(savedUser)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
