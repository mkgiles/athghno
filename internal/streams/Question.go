package streams

type QuestionAS2 struct {
	IntransitiveActivityAS2
	OneOf  *PropertyAS2 `json:"oneOf,omitempty"`
	AnyOf  *PropertyAS2 `json:"anyOf,omitempty"`
	Closed *PropertyAS2 `json:"closed,omitempty"`
}
