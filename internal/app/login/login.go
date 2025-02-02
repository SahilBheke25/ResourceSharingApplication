package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/handle"
)

type userCredentials struct {
	Username     string `json: "Username`
	Userpassword string `json: "Userpassword`
}

func Verify(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	fmt.Printf("%s", body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user userCredentials
	json.Unmarshal(body, &user)
	fmt.Printf("Username: %s, Password: %s \n", user.Username, user.Userpassword)

	handle.HandleResponse(w, user, r)

}
