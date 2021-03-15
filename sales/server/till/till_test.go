package till

import (
	"errors"
	"github.com/rogelioConsejo/till/sales/server/catalog/item"
	"math/big"
	"testing"
)

//TODO: Test w/several itemPrices, test negative prices, test with change, test paying not enough
func TestSell(t *testing.T) {
	var sellItemDouble item.Sell
	var till Sell

	//Setup the Item and the Till
	const itemPrice = 100
	sellItemDouble = setupItemDouble(itemPrice)
	till = setupTill()

	//Add the Item to the Till
	err := addItem(till, sellItemDouble)

	if thereIsError(err) {
		//Fail Test and report error
		t.Error(err.Error())
	} else if isTotalPriceCorrect(till, itemPrice) {
		err, change := pay(till, itemPrice)

		//There should be no error and change should be zero
		if isNotZero(change) {
			t.Error("change is incorrect")
		}
		if thereIsError(err) {
			t.Error(err.Error())
		}
	} else {
		t.Error("total price in till incorrect")
	}

}

func setupItemDouble(price int) item.Sell {
	return nil
}

func setupTill() Sell {
	return nil
}

func addItem(till Sell, itemDouble item.Sell) error {
	if till != nil {
		return till.Add(itemDouble)
	} else {
		return errors.New("till not created")
	}
}

func isTotalPriceCorrect(till Sell, price int) bool {
	return false
}

func pay(till Sell, itemPrice int) (error, *big.Float) {
	return till.CloseSale(big.NewFloat(float64(itemPrice)))
}

func isNotZero(change *big.Float) bool {
	return big.NewFloat(0).Cmp(change) != 0
}

func thereIsError(err error) bool {
	return err != nil
}

