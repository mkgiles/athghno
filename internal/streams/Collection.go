package streams

type CollectionAS2 struct {
	ObjectAS2
	TotalItems PropertyAS2 `json:"totalItems"`
	Current    PropertyAS2 `json:"current"`
	First      PropertyAS2 `json:"first"`
	Last       PropertyAS2 `json:"last"`
	Items      PropertyAS2 `json:"items"`
}

type OrderedCollectionAS2 struct {
	CollectionAS2
}

type CollectionPageAS2 struct {
	CollectionAS2
	PartOf   PropertyAS2 `json:"partOf"`
	Next     PropertyAS2 `json:"next"`
	Previous PropertyAS2 `json:"previous"`
}

type OrderedCollectionPageAS2 struct {
	CollectionPageAS2
	StartIndex PropertyAS2 `json:"startIndex"`
}
