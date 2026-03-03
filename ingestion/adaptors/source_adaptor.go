package adapters

type SourceAdapter interface {
	Name() string
	Fetch() ([]byte, error)
}