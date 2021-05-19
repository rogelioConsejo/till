package sales_log

type Persistence interface {
	SaleSaver
	SalesRetriever
}

type SaleSaver interface {
	Save(Sale) (id string, err error)
}

type SalesRetriever interface {
	Sales() ([]Sale, error)
	Sale(id string) (Sale, error)
}