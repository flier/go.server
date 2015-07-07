package base

type App interface {
	Run(args []string) error
}

type AppSlots struct {
	OnInit   Slot
	PreMain  Slot
	PostMain Slot
	OnExit   Slot
}

type BaseApp struct {
	Slots AppSlots

	Main func() error
}

func (a *BaseApp) Run(args []string) error {
	a.Slots.PreMain.Publish(a)

	defer a.Slots.PostMain.Publish(a)

	if a.Main != nil {
		return a.Main()
	}

	return nil
}
