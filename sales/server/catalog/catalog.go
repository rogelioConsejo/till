package catalog

import "github.com/rogelioConsejo/till/sales/server/catalog/item"

type Catalog interface {
	Navigate
	Administrate
}

type Navigate interface {
	Items() []item.Item
	GetItem(name string) item.Item
}

type Administrate interface {
	AddItem(item.Item)
	RemoveItem(item.Item)
	UpdateItem(item.Item)
}