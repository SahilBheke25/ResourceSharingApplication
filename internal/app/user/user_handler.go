package User

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	username string `json: username`
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("Error while Reading request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user user
	err = json.Unmarshal(body, &user)

	if err != nil {
		err = fmt.Errorf("Error while Unmarshelling: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// err = Repository.GetUserByUsername(user.username)

	if err != nil {
		err = fmt.Errorf("Error while Fetching user by username: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HandleResponse(w, "Success", r)
}

func HandleResponse(w http.ResponseWriter, message any, r *http.Request) {

	res, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
