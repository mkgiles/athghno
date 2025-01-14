package streams

type IntransitiveActivityAS2 struct {
	ObjectAS2
	Actor      *PropertyAS2 `json:"actor,omitempty"`
	Target     *PropertyAS2 `json:"target,omitempty"`
	Result     *PropertyAS2 `json:"result,omitempty"`
	Origin     *PropertyAS2 `json:"origin,omitempty"`
	Instrument *PropertyAS2 `json:"instrument,omitempty"`
}
