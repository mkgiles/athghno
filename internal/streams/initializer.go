package streams

func InitRegistry() *TypeRegistry {
	var registry TypeRegistry
	// Core Types
	registry.Register("Object", ObjectAS2{})
	registry.Register("Link", LinkAS2{})
	registry.Register("IntransitiveActivity", IntransitiveActivityAS2{})
	registry.Register("Activity", ActivityAS2{})
	registry.Register("Collection", CollectionAS2{})
	registry.Register("OrderedCollection", OrderedCollectionAS2{})
	registry.Register("CollectionPage", CollectionPageAS2{})
	registry.Register("OrderedCollectionPage", OrderedCollectionPageAS2{})
	// Actor Types
	registry.Register("Actor", ActorAS2{})
	registry.Register("Application", ApplicationAS2{})
	registry.Register("Group", GroupAS2{})
	registry.Register("Person", PersonAS2{})
	registry.Register("Organization", OrganizationAS2{})
	registry.Register("Service", ServiceAS2{})
	// Object Types
	registry.Register("Article", ArticleAS2{})
	registry.Register("Audio", AudioAS2{})
	registry.Register("Image", ImageAS2{})
	registry.Register("Video", VideoAS2{})
	registry.Register("Document", DocumentAS2{})
	registry.Register("Event", EventAS2{})
	registry.Register("Note", NoteAS2{})
	registry.Register("Page", PageAS2{})
	registry.Register("Place", PlaceAS2{})
	registry.Register("Profile", ProfileAS2{})
	registry.Register("Relationship", RelationshipAS2{})
	registry.Register("Tombstone", TombstoneAS2{})
	// Activity Types
	registry.Register("Accept", AcceptAS2{})
	registry.Register("Reject", RejectAS2{})
	registry.Register("Add", AddAS2{})
	registry.Register("Remove", RemoveAS2{})
	registry.Register("Update", UpdateAS2{})
	registry.Register("View", ViewAS2{})
	registry.Register("Announce", AnnounceAS2{})
	registry.Register("Arrive", ArriveAS2{})
	registry.Register("Block", BlockAS2{})
	registry.Register("Create", CreateAS2{})
	registry.Register("Delete", DeleteAS2{})
	registry.Register("Like", LikeAS2{})
	registry.Register("Dislike", DislikeAS2{})
	registry.Register("Flag", FlagAS2{})
	registry.Register("Follow", FollowAS2{})
	registry.Register("Ignore", IgnoreAS2{})
	registry.Register("Invite", InviteAS2{})
	registry.Register("Join", JoinAS2{})
	registry.Register("Leave", LeaveAS2{})
	registry.Register("Listen", ListenAS2{})
	registry.Register("Move", MoveAS2{})
	registry.Register("Offer", OfferAS2{})
	registry.Register("Question", QuestionAS2{})
	registry.Register("Read", ReadAS2{})
	registry.Register("TentativeAccept", TentativeAcceptAS2{})
	registry.Register("TentativeReject", TentativeRejectAS2{})
	registry.Register("Travel", TravelAS2{})
	registry.Register("Undo", UndoAS2{})
	return &registry
}
