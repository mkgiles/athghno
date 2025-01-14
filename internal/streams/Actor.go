package streams

/*
While the ActivityStreams Specification only expects Actor objects to inherit properties from Object
ActivityPub Mandates and Permits certain extensions to all Actor types
*/
type ActorAS2 struct {
	ObjectAS2
	Inbox                *PropertyAS2 `json:"inbox"`
	Outbox               *PropertyAS2 `json:"outbox"`
	Following            *PropertyAS2 `json:"following"`
	Followers            *PropertyAS2 `json:"followers"`
	Liked                *PropertyAS2 `json:"liked"`
	Streams              *PropertyAS2 `json:"streams,omitempty"`
	PreferredUsername    *PropertyAS2 `json:"preferredUsername,omitempty"`
	PreferredUsernameMap *PropertyAS2 `json:"preferredUsernameMap,omitempty"`
	Endpoints            *PropertyAS2 `json:"endpoints,omitempty"`
}

// Required for the Endpoints property permitted for Actor types in ActivityPub
type EndpointAS2 struct {
	ObjectAS2
	ProxyUrl                   *PropertyAS2 `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint *PropertyAS2 `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         *PropertyAS2 `json:"oauthTokenEndpoint,omitempty"`
	ProvideClientKey           *PropertyAS2 `json:"provideClientKey,omitempty"`
	SignClientKey              *PropertyAS2 `json:"signClientKey,omitempty"`
	SharedInbox                *PropertyAS2 `json:"sharedInbox,omitempty"`
}

type PersonAS2 struct {
	ActorAS2
}

type ApplicationAS2 struct {
	ActorAS2
}

type OrganizationAS2 struct {
	ActorAS2
}

type GroupAS2 struct {
	ActorAS2
}

type ServiceAS2 struct {
	ActorAS2
}
