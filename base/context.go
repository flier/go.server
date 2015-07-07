package base

import (
	"golang.org/x/net/context"
)

type Context context.Context

type ContextAware interface {
	SetContext(ctxt Context)
}
