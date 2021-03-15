package till

import (
	"github.com/rogelioConsejo/till/sales/server/catalog/item"
	"math/big"
)

type Till interface {
	Sell
	someOtherFunc()
}

type Sell interface {
	Order
	Pay
}

type Order interface {
	Add(item.Sell) error
	Remove(itemName string) error
}

type Pay interface {
	Total() *big.Float
	CloseSale(amountPayed *big.Float) (err error, change *big.Float)
}