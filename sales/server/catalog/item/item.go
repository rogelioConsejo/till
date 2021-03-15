package item

import "math/big"

type Item interface {
	Display
	Sell
	Setup
}

type DisplaySell interface {
	Display
	Sell
}

type DisplaySetup interface {
	Display
	Setup
}

type Display interface {
	Name() string
	UnitPrice() *big.Float
}

type Setup interface {
	SetName(string)
	SetUnitPrice(*big.Float)
}

type Sell interface {
	SetAmount(*big.Float)
	Amount() *big.Float
	Price() *big.Float
}
