package lion

type Lion struct {
}

func NewEngine() *Lion {
	return &Lion{}
}

func (l *Lion) Get(pattern string, handlefunc ...Handlefunc) {}
