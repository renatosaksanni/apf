package datafetching

type DataService interface {
	FetchAndSaveData(symbol string) error
}
