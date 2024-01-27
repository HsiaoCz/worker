package pkg

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitIdentity() error {
	var st time.Time
	st, err := time.Parse("2006-01-02", "2024-01-20")
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(1)
	return err
}
