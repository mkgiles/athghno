package streams

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

/*
All ActivityStreams Properties can be one of a URL, an object, or an array of the former two.
So it is necessary to wrap each property in a tagged union-like structure for flexibility.
*/
type PropertyAS2 struct {
	Simple   string        `json:",omitempty"`
	Complex  interface{}   `json:",omitempty"`
	Compound []interface{} `json:",omitempty"`
}

func (p PropertyAS2) MarshalJSON() ([]byte, error) {
	if p.IsSimple() {
		return json.Marshal(p.Simple)
	} else if p.IsComplex() {
		return json.Marshal(p.Complex)
	} else {
		return json.Marshal(p.Compound)
	}
}

func (p PropertyAS2) IsNull() bool {
	return !(p.IsSimple() && p.IsComplex() && p.IsCompound())
}
func (p PropertyAS2) Clear() {
	p.Simple = ""
	p.Complex = nil
	p.Compound = nil
}
func (p PropertyAS2) IsSimple() bool {
	return p.Simple != ""
}
func (p PropertyAS2) IsComplex() bool {
	return !p.IsSimple() && p.Complex != nil
}
func (p PropertyAS2) IsCompound() bool {
	return !p.IsSimple() && !p.IsComplex() && p.Compound != nil
}
func (p PropertyAS2) GetType() string {
	if p.IsSimple() {
		return "Simple"
	} else if p.IsComplex() {
		return "Complex"
	} else if p.IsCompound() {
		return "Compound"
	} else {
		return "null"
	}
}
func (p PropertyAS2) GetValue() interface{} {
	if p.IsSimple() {
		return p.Simple
	} else if p.IsComplex() {
		return p.Complex
	} else {
		return p.Compound
	}
}

type TypeRegistry struct {
	registry map[string]Registrable
}

type Registrable interface{}

func (t *TypeRegistry) Register(name string, r Registrable) {
	t.registry[name] = r // stores an uninitialized version of the type at the associated name
}

func (t *TypeRegistry) NewInstance(typing string) (interface{}, error) {
	if registrable, ok := t.registry[typing]; ok {
		return registrable, nil
	}
	return nil, errors.New("expected type not registered")
}

// Creates a JSON byte string from a struct
func (t *TypeRegistry) MarshalFromAS2Type(as2 interface{}) ([]byte, error) {
	var newMap = make(map[string]interface{})
	temp, err := json.Marshal(as2)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(temp, &newMap) // Populates the map with all fields
	if err != nil {
		return nil, err
	}
	value := reflect.ValueOf(as2)
	omap := value.FieldByName("Map_") // Unflatten the Map_ property to add to the map
	// Append all extra properties to map
	for k, v := range omap.Interface().(map[string]*PropertyAS2) {
		newMap[k] = v
	}
	// return json object corresponding to map
	return json.Marshal(newMap)
}

// Takes a JSON input in bytes, extracts the registered type, and unmarshals the json data into the corresponding struct
func (t *TypeRegistry) UnmarshalIntoAS2Type(jsonValue []byte) (interface{}, error) {
	// Unmarshal values into a map
	var imap map[string]interface{}
	err := json.Unmarshal(jsonValue, &imap)
	if err != nil {
		return nil, err
	}
	// Extract type parameter from map and create a reflection type and value for that type
	instance, err := t.NewInstance(imap["type"].(string))
	typing := reflect.TypeOf(instance)
	v := reflect.New(typing).Elem()
	if err != nil {
		return nil, err
	}
	// Process each field
	for _, field := range reflect.VisibleFields(typing) {
		t := field.Type
		tag := strings.Split(field.Tag.Get("json"), ",")[0]                // We grab this to map JSON field names to struct field names
		if tag != "-" && tag != "" && v.FieldByName(field.Name).CanSet() { // Grab only named properties that can be set
			// Skip fields with no set values
			if _, ok := imap[tag]; !ok {
				continue
			}
			// Create a new property
			property := toProperty(imap[tag])
			// Depending on whether the field is nullable or not, store the property itself or a pointer to it
			if t.Kind() == reflect.Ptr {
				v.FieldByName(field.Name).Set(reflect.ValueOf(&property))
				delete(imap, tag)
			} else {
				v.FieldByName(field.Name).Set(reflect.ValueOf(property))
				delete(imap, tag)
			}
		}
	}
	// Make a new Map_ property for all remaining fields
	nmap := make(map[string]*PropertyAS2, len(imap))
	for key, val := range imap {
		newValue := toProperty(val)
		nmap[key] = &newValue
	}
	v.FieldByName("Map_").Set(reflect.ValueOf(nmap))
	return v.Interface(), nil
}

// Takes an ordinary input and wraps it in a Property
func toProperty(m interface{}) PropertyAS2 {
	kind := reflect.TypeOf(m).Kind()
	if kind == reflect.String {
		return PropertyAS2{Simple: m.(string)}
	} else if kind == reflect.Ptr {
		return *m.(*PropertyAS2)
	} else if kind == reflect.Array {
		return PropertyAS2{Compound: m.([]interface{})}
	} else {
		return PropertyAS2{Complex: m.(interface{})}
	}
}
