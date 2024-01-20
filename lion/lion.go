package lion

type Lion struct {
}

func NewEngine() *Lion {
	return &Lion{}
}

func (l *Lion) Get(pattern string, handlefunc ...Handlefunc)    {}
func (l *Lion) Post(pattern string, handlefunc ...Handlefunc)   {}
func (l *Lion) Delete(pattern string, handlefunc ...Handlefunc) {}
func (l *Lion) Put(pattern string, handlefunc ...Handlefunc)    {}
func (l *Lion) Trance(pattern string, handlefunc ...Handlefunc) {}
