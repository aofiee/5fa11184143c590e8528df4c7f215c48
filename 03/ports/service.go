package ports

import "count-beef/domain"

type Service interface {
	GetSummary() (domain.Response, error)
}
