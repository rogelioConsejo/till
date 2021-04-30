package sales_log

type Persistence interface {
	Save(Sale) error
	Sales() ([]Sale, error)
}