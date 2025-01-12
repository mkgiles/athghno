package streams

type TombstoneAS2 struct {
	ObjectAS2
	FormerType PropertyAS2 `json:"formerType"`
	Deleted    PropertyAS2 `json:"deleted"`
}
