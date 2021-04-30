package sales_log

type Sale struct{
	items []Item
}

type Item struct {
	price int
	name string
}

func NewSale(items []Item) *Sale {
	sale := new(Sale)
	sale.items = items

	return sale
}
