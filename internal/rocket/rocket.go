//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/mauricewittek/go-grpc-service/internal/rocket Store

package rocket

import (
	"context"
)

type Rocket struct {
	ID   string
	Name string
	Type string
}

type Store interface {
	GetRocket(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetRocket(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocket(id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}

	return nil
}
