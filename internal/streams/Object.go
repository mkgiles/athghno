package streams

type ObjectAS2 struct {
	Registrable                         // Makes all objects that embed ObjectAS2 registrable in the TypeRegistry
	LDContext    PropertyAS2            `json:"@context"`
	Id           PropertyAS2            `json:"id"`
	Type         PropertyAS2            `json:"type"`
	Name         PropertyAS2            `json:"name"`
	NameMap      PropertyAS2            `json:"nameMap"`
	Summary      PropertyAS2            `json:"summary"`
	SummaryMap   PropertyAS2            `json:"summaryMap"`
	Icon         PropertyAS2            `json:"icon"`
	Image        PropertyAS2            `json:"image"`
	Tag          PropertyAS2            `json:"tag"`
	Content      PropertyAS2            `json:"content"`
	ContentMap   PropertyAS2            `json:"contentMap"`
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
	Source       PropertyAS2            `json:"source"` // ActivityPub Extension of ActivityStreams Property
	Likes        PropertyAS2            `json:"likes"`  // ditto
	Shares       PropertyAS2            `json:"shares"` // ^^
	Map_         map[string]PropertyAS2 `json:"-"`      // Contains all properties that are not in spec
}
