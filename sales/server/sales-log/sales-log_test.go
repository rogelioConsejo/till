package sales_log

import (
	"errors"
	"testing"
)


type MockPersistence struct {
	data []Sale
}

func (p *MockPersistence) Save(sale Sale) error {
	p.data = append(p.data, sale)
	return nil
}

func (p MockPersistence) Sales() (sales []Sale, err error) {
	return p.data, nil
}

func TestLogSale(t *testing.T) {
	const price1 = 1
	const name1 = "product 1"
	const price2 = 2
	const name2 = "product 2"

	items := createMockItems(price1, name1, price2, name2)

	salesLogger := SalesLog(new(MockPersistence))
	sale := NewSale(items)

	err := salesLogger.Log(sale)
	if err != nil {
		t.Error(err)
	}

	sales, err := salesLogger.persistence.Sales()

	if err != nil {
		t.Error(err)
	} else if len(sales) != 1 {
		t.Error(errors.New("sales do not correspond to logged sale"))
	} else {
		if len(sales[0].items) != 2 {
			t.Error(errors.New("items do not correspond to logged sale"))
		}
	}
}

func createMockItems(price1 int, name1 string, price2 int, name2 string) []Item {
	item1 := new(Item)
	item1.name = name1
	item1.price = price1
	item2 := new(Item)
	item1.name = name2
	item1.price = price2
	items := []Item{*item1, *item2}
	return items
}