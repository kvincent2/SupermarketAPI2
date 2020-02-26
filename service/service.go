package service

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/SupermarketAPI2/db"
)

type Service struct {
	Database *db.Database
	Router 	 *mux.Router
}

func New(logger log.Logger, db *db.Database) {
	service := Service{
		Database: nil,
		Router:   mux.NewRouter(),
	}

}

func (s *Service) setupGetAllHandler () {
	getAllEndpoint := newGetAllEndpoint()
}

func newGetAllEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.GetAll(ctx)
	}
}
