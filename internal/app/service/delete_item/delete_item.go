package delete_item

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sahirrrr/PARI-Test/internal/app/repo"
	"go.uber.org/dig"
)

type (
	DeleteItemSvc interface {
		DeleteItem(
			/*req*/ ctx context.Context, request DeleteItemRequest) (
			/*res*/ response DeleteItemResponse, httpCode int, err error,
		)
	}

	DeleteItemRequest struct {
		ID int64 `json:"id"`
	}

	DeleteItemResponse struct {
		Message string `json:"message"`
	}

	Dependencies struct {
		dig.In
		PostgreSQL *sqlx.DB
	}

	Service struct {
		Dependencies
		Repo repo.PostgreSQL
	}
)

var (
	ErrNegatifID  = errors.New("error id cannot be negatif")
	ErrIDNotFound = errors.New("error id cannot be found")
)

func New(impl Dependencies) DeleteItemSvc {
	return &Service{Dependencies: impl, Repo: repo.NewPostgreSQL(impl.PostgreSQL)}
}

func (x *Service) DeleteItem(
	/*req*/ ctx context.Context, request DeleteItemRequest) (
	/*res*/ response DeleteItemResponse, httpCode int, err error,
) {
	if request.ID <= 0 {
		return response, http.StatusBadRequest, ErrNegatifID
	}

	item, err := x.Repo.GetItemByID(ctx, repo.GetItemByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		log.Printf("[DeleteItem][GetItemByID] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if item.ID == 0 {
		return response, http.StatusNotFound, ErrIDNotFound
	}

	_, err = x.Repo.DeleteItem(ctx, repo.DeleteItemRequest{
		ID: request.ID,
	})
	if err != nil {
		log.Printf("[DeleteItem][[DeleteItem]] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	response.Message = fmt.Sprintf("Item with ID : %d, has been deleted", request.ID)

	return response, http.StatusOK, nil
}
