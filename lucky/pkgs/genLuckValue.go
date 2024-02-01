package pkgs

import (
	"math/rand"
	"time"
)

func GenLuckValue() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
}
