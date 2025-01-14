package streams

type LinkAS2 struct {
	Registrable `json:"-"`              // Makes all objects that embed LinkAS2 registrable in the TypeRegistry
	Context     PropertyAS2             `json:"@context"`
	Id          PropertyAS2             `json:"id"`
	Type        PropertyAS2             `json:"type"`
	Href        *PropertyAS2            `json:"href,omitempty"`
	HrefLang    *PropertyAS2            `json:"hreflang,omitempty"`
	Rel         *PropertyAS2            `json:"rel,omitempty"`
	Name        *PropertyAS2            `json:"name,omitempty"`
	Preview     *PropertyAS2            `json:"preview,omitempty"`
	MediaType   *PropertyAS2            `json:"mediaType,omitempty"`
	Height      *PropertyAS2            `json:"height,omitempty"`
	Width       *PropertyAS2            `json:"width,omitempty"`
	Map_        map[string]*PropertyAS2 `json:"-"`
}
