package catalog

import "github.com/rogelioConsejo/till/item"

type Catalog interface {
	Navigator
	Administrator
}

type Navigator interface {
	Items() []item.Item
	GetItem(name string) item.Item
}

type Administrator interface {
	AddItem(item.Item)
	RemoveItem(item.Item)
	UpdateItem(item.Item)
}