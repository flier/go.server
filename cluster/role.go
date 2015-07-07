package cluster

type Role int

const (
	JoinElection Role = iota
	AsLeader
	AsFollower
	AsWatcher
)
