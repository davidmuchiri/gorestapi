package utils

import (
	"encoding/json"
	"net/http"
)

//Message is a function that returns a json response
func Message(status bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message, "data": data}
}

// Response returns an actual http response
func Response(res http.ResponseWriter, data map[string]interface{}) {
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}
