package types

type FriendshipStatus string

var (
	Pending  FriendshipStatus = "pending"
	Accepted FriendshipStatus = "accepted"
	Rejected FriendshipStatus = "rejected"
)
