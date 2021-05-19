package sales_log

type Persistence interface {
	SaleSaver
	SalesRetriever
}

type SaleSaver interface {
	Save(Sale) (id string, err error)
}

type SalesRetriever interface {
	MultipleSalesRetriever
	SingleSaleRetriever
}

type MultipleSalesRetriever interface {
	Sales() ([]Sale, error)
}

type SingleSaleRetriever interface {
	Sale(id string) (Sale, error)
}