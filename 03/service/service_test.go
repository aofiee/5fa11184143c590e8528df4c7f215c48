package service

import (
	mockrepository "count-beef/testings/mock_repository"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

type SetupTest struct {
	repo *mockrepository.MockRepo
	srv  *Service
}

func Setup(t *testing.T) *SetupTest {
	ctrl := gomock.NewController(t)
	repo := mockrepository.NewMockRepo(ctrl)
	srv := NewService(repo)
	return &SetupTest{
		repo: repo,
		srv:  srv,
	}
}

func TestNewService(t *testing.T) {
	Setup(t)
}

func TestGetSummary(t *testing.T) {
	st := Setup(t)

	t.Run("Success", func(t *testing.T) {
		st.repo.EXPECT().GetSummary().Return("t-bone beef t-bone", nil)
		rs, err := st.srv.GetSummary()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		assert.Equal(t, rs.Beef["t-bone"], 2)
	})

	t.Run("Failed", func(t *testing.T) {
		st.repo.EXPECT().GetSummary().Return("", errors.New("Failed to fetch data"))
		_, err := st.srv.GetSummary()
		if err == nil {
			t.Errorf("Error: %v", err)
		}
	})
}
