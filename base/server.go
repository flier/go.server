package base

type Server interface {
	Service
}

type BaseServer struct {
	*BaseService
}

func NewBaseServer(name string) *BaseServer {
	return &BaseServer{NewBaseService(name)}
}
