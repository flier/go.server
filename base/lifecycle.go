package base

type AppLifeCycleAware interface {
	OnInit(app App)

	PreMain(app App)

	PostMain(app App)

	OnExit(app App)
}

type ServiceLifecycleAware interface {
	OnStart(svc Service)

	PreServe(svc Service)

	PostServe(svc Service)

	OnStop(svc Service)
}
