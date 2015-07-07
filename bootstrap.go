package server

import (
	"github.com/flier/go.server/base"
	"github.com/flier/go.server/cluster"
)

type Bootstrap struct {
	base    *base.Bootstrap
	cluster *cluster.Bootstrap
}

func NewBootstrap(name string) *Bootstrap {
	return &Bootstrap{base: base.NewBootstrap(name)}
}

func (b *Bootstrap) Base() *base.Bootstrap {
	if b.base == nil {
		b.base = &base.Bootstrap{}
	}

	return b.base
}

func (b *Bootstrap) Cluster() *cluster.Bootstrap {
	if b.cluster == nil {
		b.cluster = cluster.NewBootstrap()
	}

	return b.cluster
}

func (b *Bootstrap) App() base.App {
	return nil
}
