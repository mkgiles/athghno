package streams

type IntransitiveActivityAS2 struct {
	ObjectAS2
	Actor      PropertyAS2 `json:"actor"`
	Target     PropertyAS2 `json:"target"`
	Result     PropertyAS2 `json:"result"`
	Origin     PropertyAS2 `json:"origin"`
	Instrument PropertyAS2 `json:"instrument"`
}
