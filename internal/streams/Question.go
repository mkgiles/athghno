package streams

type QuestionAS2 struct {
	IntransitiveActivityAS2
	OneOf  PropertyAS2 `json:"oneOf"`
	AnyOf  PropertyAS2 `json:"anyOf"`
	Closed PropertyAS2 `json:"closed"`
}
