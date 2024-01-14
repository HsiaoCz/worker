package data

type BookTicker struct {
	UpdateId      int64  `json:"u"` // update id
	Symbol        string `json:"s"` // 交易对
	BuyBestPrice  string `json:"b"` // 买单最优挂单价格
	BuyBestCount  string `json:"B"` // 买单最优挂单数量
	SealBestPrice string `json:"a"` // 卖单最优挂单价格
	SealBestCount string `json:"A"` // 卖单最优挂单数据
}
