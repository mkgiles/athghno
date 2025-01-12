package streams

type LinkAS2 struct {
	Context   PropertyAS2   `json:"@context"`
	Id        PropertyAS2   `json:"id"`
	Type      PropertyAS2   `json:"type"`
	Href      PropertyAS2   `json:"href"`
	HrefLang  PropertyAS2   `json:"hreflang"`
	Rel       PropertyAS2   `json:"rel"`
	Name      PropertyAS2   `json:"name"`
	Preview   PropertyAS2   `json:"preview"`
	MediaType PropertyAS2   `json:"mediaType"`
	Height    PropertyAS2   `json:"height"`
	Width     PropertyAS2   `json:"width"`
	Map_      []PropertyAS2 `json:"-"`
}
