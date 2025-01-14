package streams

type PlaceAS2 struct {
	ObjectAS2
	Accuracy  *PropertyAS2 `json:"accuracy,omitempty"`
	Altitude  *PropertyAS2 `json:"altitude,omitempty"`
	Latitude  *PropertyAS2 `json:"latitude,omitempty"`
	Longitude *PropertyAS2 `json:"longitude,omitempty"`
	Radius    *PropertyAS2 `json:"radius,omitempty"`
	Units     *PropertyAS2 `json:"units,omitempty"`
}
