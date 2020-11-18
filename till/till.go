package till

import (
	"github.com/rogelioConsejo/till/item"
	"math/big"
)

type Till interface {
	Order
	Payment
}

type Order interface {
	Add(item.Item)
	Remove(item.Item)
}

type Payment interface {
	Total() big.Float
	CloseSale(amountPayed big.Float) (change big.Float)
}