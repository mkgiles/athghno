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

func (t *TypeRegistry) UnmarshalIntoAS2Type(jsonValue []byte) (interface{}, error) {
	var imap map[string]interface{}
	err := json.Unmarshal(jsonValue, &imap)
	if err != nil {
		return nil, err
	}
	instance, err := t.NewInstance(imap["type"].(string))
	typing := reflect.TypeOf(instance)
	v := reflect.New(typing).Elem()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonValue, &imap)
	if err != nil {
		return nil, err
	}
	for _, field := range reflect.VisibleFields(typing) {
		t := field.Type
		tag := strings.Split(field.Tag.Get("json"), ",")[0]
		if tag != "-" && tag != "" && v.FieldByName(field.Name).CanSet() {
			if _, ok := imap[tag]; !ok {
				continue
			}
			property := PropertyAS2{}
			propType := reflect.TypeOf(imap[tag]).Kind()
			if propType == reflect.String {
				property.Simple = imap[tag].(string)
			} else if propType == reflect.Array {
				property.Compound = imap[tag].([]interface{})
			} else {
				property.Complex = imap[tag]
			}
			if t.Kind() == reflect.Ptr {
				v.FieldByName(field.Name).Set(reflect.ValueOf(&property))
			} else {
				v.FieldByName(field.Name).Set(reflect.ValueOf(property))
			}
		}
	}
	return v.Interface(), nil
}
