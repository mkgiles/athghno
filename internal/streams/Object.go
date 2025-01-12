package streams

type ObjectAS2 struct {
	LDContext    PropertyAS2            `json:"@context"`
	Id           PropertyAS2            `json:"id"`
	Type         PropertyAS2            `json:"type"`
	Name         PropertyAS2            `json:"name"`
	Summary      PropertyAS2            `json:"summary"`
	Icon         PropertyAS2            `json:"icon"`
	Image        PropertyAS2            `json:"image"`
	Tag          PropertyAS2            `json:"tag"`
	Content      PropertyAS2            `json:"content"`
	Published    PropertyAS2            `json:"published"`
	Updated      PropertyAS2            `json:"updated"`
	Url          PropertyAS2            `json:"url"`
	AttributedTo PropertyAS2            `json:"attributedTo"`
	To           PropertyAS2            `json:"to"`
	BTo          PropertyAS2            `json:"bto"`
	CcTo         PropertyAS2            `json:"cc"`
	Bcc          PropertyAS2            `json:"bcc"`
	InReplyTo    PropertyAS2            `json:"inReplyTo"`
	Attachment   PropertyAS2            `json:"attachment"`
	Audience     PropertyAS2            `json:"audience"`
	Replies      PropertyAS2            `json:"replies"`
	Location     PropertyAS2            `json:"location"`
	Preview      PropertyAS2            `json:"preview"`
	Generator    PropertyAS2            `json:"generator"`
	MediaType    PropertyAS2            `json:"mediaType"`
	StartTime    PropertyAS2            `json:"startTime"`
	EndTime      PropertyAS2            `json:"endTime"`
	Duration     PropertyAS2            `json:"duration"`
	Context      PropertyAS2            `json:"context"`
	Map_         map[string]PropertyAS2 `json:"-"`
}
