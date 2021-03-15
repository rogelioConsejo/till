package server

import (
	"github.com/rogelioConsejo/till/sales/server/catalog/item"
	"math/big"
)

type Sale interface {
	Id() string
	Total() *big.Float
	Close(amountPaid *big.Float) (change *big.Float)
}

type sale struct {
	id string
	items []item.Sell
}

func (s *sale) Id() string {
	return s.id
}

//TODO
func (s *sale) Total() *big.Float{
	return big.NewFloat(0)
}

//TODO
func (s *sale) Close(amountPaid *big.Float) (change *big.Float){
	return big.NewFloat(0)
}