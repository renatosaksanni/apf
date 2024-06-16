package datafetching

type DataRepository interface {
	FetchData(symbol string) ([]Data, error)
}
