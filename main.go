package main

import (
	"athghno/internal/streams"
	"fmt"
	"time"
)

func main() {
	registry := streams.InitRegistry()
	var create streams.CreateAS2
	var note streams.NoteAS2
	LDContext := streams.PropertyAS2{
		Simple:   "https://w3.org/ns/activitystreams",
		Complex:  nil,
		Compound: nil,
	}
	create.LDContext = LDContext
	create.Id.Simple = "https://example.com/create/1"
	create.Type.Simple = "Create"
	create.Actor = &streams.PropertyAS2{}
	create.Actor.Simple = "https://example.com/actor"
	note.LDContext = LDContext
	note.Id.Simple = "https://example.com/note/1"
	note.Type.Simple = "Note"
	note.Published = &streams.PropertyAS2{}
	note.Published.Simple = time.Now().Format(time.RFC3339)
	note.AttributedTo = &streams.PropertyAS2{}
	note.AttributedTo.Simple = create.Actor.Simple
	note.Content = &streams.PropertyAS2{}
	note.Content.Simple = "Hello, world!"
	note.To = &streams.PropertyAS2{}
	note.To.Simple = "https://w3.org/ns/activitystreams#Public"
	create.Object = streams.PropertyAS2{}
	create.Object.Complex = note
	create.Map_ = map[string]*streams.PropertyAS2{}
	create.Map_["test"] = &streams.PropertyAS2{Simple: "test"}
	jsonData, err := registry.MarshalFromAS2Type(create)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	result, err := registry.UnmarshalIntoAS2Type(jsonData)
	if err != nil {
		panic(err)
	}
	jsonData, err = registry.MarshalFromAS2Type(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
