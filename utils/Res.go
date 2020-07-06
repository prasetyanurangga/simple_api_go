package utils

import(
	"net/http"
	"encoding/json"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int){
	jsonEnkoding, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")

	if err != nil{
		http.Error(w, "Error", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonEnkoding))
}