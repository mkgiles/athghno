package serverProto

import (
	"athghno/internal/dataStore"
	"athghno/internal/streams"
	"encoding/json"
	"github.com/dgraph-io/badger/v4"
	"net/http"
	"strings"
)

func CreateWF(db *badger.DB, subject string, actor interface{}) error {
	var id string
	resource := strings.Split(strings.TrimPrefix(subject, "acct:"), "@")
	user := resource[0]
	domain := resource[1]
	address := domain + "/users/" + user
	switch actor.(type) {
	case streams.PersonAS2:
		person := actor.(streams.PersonAS2)
		id = person.Id.Simple
	case streams.OrganizationAS2:
		organization := actor.(streams.OrganizationAS2)
		id = organization.Id.Simple
	case streams.ApplicationAS2:
		application := actor.(streams.ApplicationAS2)
		id = application.Id.Simple
	case streams.GroupAS2:
		group := actor.(streams.GroupAS2)
		id = group.Id.Simple
	case streams.ServiceAS2:
		service := actor.(streams.ServiceAS2)
		id = service.Id.Simple
	case streams.ActorAS2:
		as2 := actor.(streams.ActorAS2)
		id = as2.Id.Simple
	}
	t := "application/activity+json"
	addLink := WFLink{
		Rel:  "self",
		Href: &address,
		Type: &t,
	}
	finger := Finger{
		Subject: subject,
		Links:   []WFLink{addLink},
		Aliases: []string{id, address},
	}
	jsonData, err := json.Marshal(finger)
	if err != nil {
		return err
	}
	err = dataStore.PutObject(db, []byte(subject), jsonData)
	if err != nil {
		return err
	}
	return nil
}

func WebFinger(db *badger.DB, w http.ResponseWriter, r *http.Request) {
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
	domain := resource[1]
	if domain != r.Host {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid domain in resource parameter"))
	}
	jsonData, err := dataStore.GetObject(db, []byte(query.Get("resource")))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type WFLink struct {
	Rel        string            `json:"rel"`
	Type       *string           `json:"type,omitempty"`
	Href       *string           `json:"href,omitempty"`
	Titles     map[string]string `json:"titles,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}
type Finger struct {
	Subject    string            `json:"subject"`
	Aliases    []string          `json:"aliases,omitempty"`
	Links      []WFLink          `json:"links,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}
