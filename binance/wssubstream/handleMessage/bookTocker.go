package handlemessage

import (
	"fmt"

	"github.com/HsiaoCz/worker/binance/wssubstream/data"
)

func HandleBookTicker(data *data.BookTicker) {
	fmt.Println(data)
}
