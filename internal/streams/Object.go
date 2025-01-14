package streams

type ObjectAS2 struct {
	Registrable  `json:"-"`              // Makes all objects that embed ObjectAS2 registrable in the TypeRegistry
	LDContext    PropertyAS2             `json:"@context"`
	Id           PropertyAS2             `json:"id"`
	Type         PropertyAS2             `json:"type"`
	Name         *PropertyAS2            `json:"name,omitempty"`
	NameMap      *PropertyAS2            `json:"nameMap,omitempty"`
	Summary      *PropertyAS2            `json:"summary,omitempty"`
	SummaryMap   *PropertyAS2            `json:"summaryMap,omitempty"`
	Icon         *PropertyAS2            `json:"icon,omitempty"`
	Image        *PropertyAS2            `json:"image,omitempty"`
	Tag          *PropertyAS2            `json:"tag,omitempty"`
	Content      *PropertyAS2            `json:"content,omitempty"`
	ContentMap   *PropertyAS2            `json:"contentMap,omitempty"`
	Published    *PropertyAS2            `json:"published,omitempty"`
	Updated      *PropertyAS2            `json:"updated,omitempty"`
	Url          *PropertyAS2            `json:"url,omitempty"`
	AttributedTo *PropertyAS2            `json:"attributedTo,omitempty"`
	To           *PropertyAS2            `json:"to,omitempty"`
	BTo          *PropertyAS2            `json:"bto,omitempty"`
	CcTo         *PropertyAS2            `json:"cc,omitempty"`
	Bcc          *PropertyAS2            `json:"bcc,omitempty"`
	InReplyTo    *PropertyAS2            `json:"inReplyTo,omitempty"`
	Attachment   *PropertyAS2            `json:"attachment,omitempty"`
	Audience     *PropertyAS2            `json:"audience,omitempty"`
	Replies      *PropertyAS2            `json:"replies,omitempty"`
	Location     *PropertyAS2            `json:"location,omitempty"`
	Preview      *PropertyAS2            `json:"preview,omitempty"`
	Generator    *PropertyAS2            `json:"generator,omitempty"`
	MediaType    *PropertyAS2            `json:"mediaType,omitempty"`
	StartTime    *PropertyAS2            `json:"startTime,omitempty"`
	EndTime      *PropertyAS2            `json:"endTime,omitempty"`
	Duration     *PropertyAS2            `json:"duration,omitempty"`
	Context      *PropertyAS2            `json:"context,omitempty"`
	Source       *PropertyAS2            `json:"source,omitempty"` // ActivityPub Extension of ActivityStreams Property
	Likes        *PropertyAS2            `json:"likes,omitempty"`  // ditto
	Shares       *PropertyAS2            `json:"shares,omitempty"` // ^^
	Map_         map[string]*PropertyAS2 `json:"-"`                // Contains all properties that are not in spec
}
