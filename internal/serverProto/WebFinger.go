package serverProto

import (
	"encoding/json"
	"net/http"
	"strings"
)

func WebFinger(w http.ResponseWriter, r *http.Request) {
	//TODO: Connect to data store
	query := r.URL.Query()
	if len(query) < 1 || !query.Has("resource") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("resource parameter is missing"))
		return
	}
	resource := strings.Split(query.Get("resource"), "@")
	if len(resource) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid resource parameter"))
		return
	}
	user := resource[0]
	domain := resource[1]
	if domain != r.Host {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid domain in resource parameter"))
	}
	response := make(map[string]interface{})
	response["subject"] = resource
	response["links"] = []map[string]interface{}{}
	self := make(map[string]interface{})
	self["rel"] = "self"
	self["href"] = domain + "/users/" + user
	self["type"] = "application/activity+json"
	response["links"] = append(response["links"].([]map[string]interface{}), self)
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal response"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/jrd+json")
	w.Write(jsonData)
}
