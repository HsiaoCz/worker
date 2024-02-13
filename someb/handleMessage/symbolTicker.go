package handlemessage

import (
	"fmt"

	"github.com/HsiaoCz/worker/someb/data"
)

func HandleSymbolTicker(data *data.BookTicker) {
	fmt.Println(data)
}
