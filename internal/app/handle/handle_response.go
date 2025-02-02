package handle

import (
	"encoding/json"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, message any, r *http.Request) {

	res, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
