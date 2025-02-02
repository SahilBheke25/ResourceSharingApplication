package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/Repository"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/handle"
)

type userCredentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func Verify(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var user userCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = fmt.Errorf("Error while decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	verified, err := Repository.AuthenticateUser(user.Username, user.Password)
	if err != nil || !verified {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handle.HandleResponse(w, "User Verifed Successfully", r)

}
