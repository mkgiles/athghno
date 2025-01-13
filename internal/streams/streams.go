package streams

import "errors"

/*
All ActivityStreams Properties can be one of a URL, an object, or an array of the former two.
So it is necessary to wrap each property in a tagged union-like structure for flexibility.
*/
type PropertyAS2 struct {
	simple   string
	complex  interface{}
	compound []interface{}
}

func (p PropertyAS2) Simple() string {
	return p.simple
}
func (p PropertyAS2) Complex() interface{} {
	return p.complex
}
func (p PropertyAS2) Compound() interface{} {
	return p.compound
}
func (p PropertyAS2) IsNull() bool {
	return !(p.IsSimple() && p.IsComplex() && p.IsCompound())
}
func (p PropertyAS2) Clear() {
	p.simple = ""
	p.complex = nil
	p.compound = nil
}
func (p PropertyAS2) IsSimple() bool {
	return p.simple != ""
}
func (p PropertyAS2) IsComplex() bool {
	return !p.IsSimple() && p.complex != nil
}
func (p PropertyAS2) IsCompound() bool {
	return !p.IsSimple() && !p.IsComplex() && p.compound != nil
}
func (p PropertyAS2) GetType() string {
	if p.IsSimple() {
		return "simple"
	} else if p.IsComplex() {
		return "complex"
	} else if p.IsCompound() {
		return "compound"
	} else {
		return "null"
	}
}
func (p PropertyAS2) GetValue() interface{} {
	if p.IsSimple() {
		return p.simple
	} else if p.IsComplex() {
		return p.complex
	} else {
		return p.compound
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
