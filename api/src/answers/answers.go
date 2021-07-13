package answers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Return a json response
func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Err(rw http.ResponseWriter, statusCode int, err error) {
	JSON(rw, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
