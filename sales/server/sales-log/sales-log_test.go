package sales_log

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

const price1 = 1
const name1 = "product 1"
const price2 = 2
const name2 = "product 2"

type MockPersistence struct {
	data []Sale
}
func (p *MockPersistence) Save(sale Sale) (id string, err error) {
	p.data = append(p.data, sale)
	return fmt.Sprintf("%d",len(p.data)-1), nil
}
func (p *MockPersistence) Sales() (sales []Sale, err error) {
	return p.data, nil
}
func (p *MockPersistence) Sale(id string) (sale Sale, err error)  {
	idInt, err := strconv.ParseInt(id, 10, 32)

	if int(idInt) >= len(p.data) || idInt < 0 {
		err = errors.New("sale not found")
	} else {
		sale = p.data[int(idInt)]
	}
	return sale, err
}

func TestSalesLogger(t *testing.T) {
	salesLogger, saleId, err := testLogSale(t)

	if noError(err) {
		err = testRetrieveSales(t, salesLogger)
	}
	if noError(err) {
		testRetrieveOneSale(t, salesLogger, saleId)
	}
}

func testRetrieveOneSale(t *testing.T, salesLogger *Logger, id string) {
	t.Log("Testing: retrieve one sale")
	sale, err := salesLogger.Sale(id)
	if errorExists(err) {
		t.Errorf("error retrieving one sale: %s", err.Error())
	} else {
		t.Log(fmt.Sprintf("Retrieved sale: %+v", sale))
	}
	return
}
func testRetrieveSales(t *testing.T, salesLogger *Logger) error {
	t.Log("Testing: retrieve sales")
	sales, err := salesLogger.Sales()
	if err != nil {
		t.Error(err)
	} else if len(sales) != 1 {
		t.Error(errors.New("sales do not correspond to logged sale"))
	} else {
		if len(sales[0].items) != 2 {
			t.Error(errors.New("mockItems do not correspond to logged sale"))
		}
	}
	return err
}
func testLogSale(t *testing.T) (salesLogger *Logger, id string, err error) {
	mockItems := createMockItems()
	mockSalesLogger := SalesLog(new(MockPersistence))

	t.Log("Testing: logging a sale")
	id, err = logNewSale(mockItems, mockSalesLogger)
	if err != nil {
		t.Error(err)
	}
	return mockSalesLogger, id, err
}

func noError(err error) bool {
	return err == nil
}
func errorExists(err error) bool {
	return err != nil
}

func logNewSale(items []Item, salesLogger *Logger) (id string, err error) {
	sale := NewSale(items)
	return salesLogger.Log(sale)
}
func createMockItems() []Item {
	item1 := new(Item)
	item1.name = name1
	item1.price = price1
	item2 := new(Item)
	item1.name = name2
	item1.price = price2
	items := []Item{*item1, *item2}
	return items
}