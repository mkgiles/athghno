package streams

type PlaceAS2 struct {
	ObjectAS2
	Accuracy  PropertyAS2 `json:"accuracy"`
	Altitude  PropertyAS2 `json:"altitude"`
	Latitude  PropertyAS2 `json:"latitude"`
	Longitude PropertyAS2 `json:"longitude"`
	Radius    PropertyAS2 `json:"radius"`
	Units     PropertyAS2 `json:"units"`
}
