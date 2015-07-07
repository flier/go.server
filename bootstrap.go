package server

import (
	"github.com/flier/go.server/cluster"
)

type ServerBootstrap struct {
	name string
}

func Bootstrap(name string) *ServerBootstrap {
	return &ServerBootstrap{
		name: name,
	}
}

func (b *ServerBootstrap) ClusterRole(role cluster.Role) *ServerBootstrap {

}

func (b *ServerBootstrap) AsLeader() *ServerBootstrap {

}

func (b *ServerBootstrap) AsLeader() *ServerBootstrap {

}

func (b *ServerBootstrap) Build() Server {

}
