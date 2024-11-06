package service

import (
	"count-beef/domain"
	"count-beef/ports"
	"log"
	"regexp"
)

type (
	Service struct {
		Repo ports.Repo
	}
)

func NewService(repo ports.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetSummary() (domain.Response, error) {
	var (
		res = domain.Response{Beef: make(map[string]int)}
	)
	txt, err := s.Repo.GetSummary()
	if err != nil {
		return res, err
	}

	for _, target := range domain.Target {
		re := regexp.MustCompile(`\b` + target + `\b`)
		matches := re.FindAllString(txt, -1)
		log.Println(target, "count:", len(matches))
		res.Beef[target] = len(matches)
	}
	return res, nil
}
