package streams

type RelationshipAS2 struct {
	ObjectAS2
	Subject      PropertyAS2 `json:"subject"`
	Object       PropertyAS2 `json:"object"`
	Relationship PropertyAS2 `json:"relationship"`
}
