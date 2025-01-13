package streams

/*
While the ActivityStreams Specification only expects Actor objects to inherit properties from Object
ActivityPub Mandates and Permits certain extensions to all Actor types
*/
type ActorAS2 struct {
	ObjectAS2
	Inbox                PropertyAS2 `json:"inbox"`
	Outbox               PropertyAS2 `json:"outbox"`
	Following            PropertyAS2 `json:"following"`
	Followers            PropertyAS2 `json:"followers"`
	Liked                PropertyAS2 `json:"liked"`
	Streams              PropertyAS2 `json:"streams"`
	PreferredUsername    PropertyAS2 `json:"preferredUsername"`
	PreferredUsernameMap PropertyAS2 `json:"preferredUsernameMap"`
	Endpoints            PropertyAS2 `json:"endpoints"`
}

// Required for the Endpoints property permitted for Actor types in ActivityPub
type EndpointAS2 struct {
	ObjectAS2
	ProxyUrl                   PropertyAS2 `json:"proxyUrl"`
	OauthAuthorizationEndpoint PropertyAS2 `json:"oauthAuthorizationEndpoint"`
	OauthTokenEndpoint         PropertyAS2 `json:"oauthTokenEndpoint"`
	ProvideClientKey           PropertyAS2 `json:"provideClientKey"`
	SignClientKey              PropertyAS2 `json:"signClientKey"`
	SharedInbox                PropertyAS2 `json:"sharedInbox"`
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
