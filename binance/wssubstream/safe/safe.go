package safe

import "sync"

var (
	Wg   sync.WaitGroup
	L    sync.Mutex
	RWL  sync.RWMutex
	SMap sync.Map
)
