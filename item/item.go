package item

import "math/big"

type Item interface {
	Display
	Config
	Sale
}

type Display interface {
	Name() string
	UnitPrice() big.Float
}

type Config interface {
	SetName(string)
	SetUnitPrice(big.Float)
}

type Sale interface {
	SetAmount(big.Float)
	Amount() big.Float
}
