package ports

type Repo interface {
	GetSummary() (string, error)
}
