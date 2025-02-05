package get_item_by_id

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
	GetItemByIDSvc interface {
		GetItemByID(
			/*req*/ ctx context.Context, request GetItemByIDRequest) (
			/*res*/ response GetItemByIDResponse, httpCode int, err error,
		)
	}

	GetItemByIDRequest struct {
		ID int64 `json:"id"`
	}

	GetItemByIDResponse struct {
		ID   int64       `json:"id"`
		Name string      `json:"name"`
		Data ItemDetails `json:"data"`
	}

	ItemDetails struct {
		ItemDetailsID int64    `json:"item_details_id"`
		CPUModel      *string  `json:"cpu_model,omitempty"`
		RAM           *string  `json:"ram,omitempty"`
		Year          *int64   `json:"year,omitempty"`
		ScreenSize    *string  `json:"screen_size,omitempty"`
		Capacity      *string  `json:"capacity,omitempty"`
		Color         *string  `json:"color,omitempty"`
		Price         *float64 `json:"price,omitempty"`
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

func New(impl Dependencies) GetItemByIDSvc {
	return &Service{Dependencies: impl, Repo: repo.NewPostgreSQL(impl.PostgreSQL)}
}

func (x *Service) GetItemByID(
	/*req*/ ctx context.Context, request GetItemByIDRequest) (
	/*res*/ response GetItemByIDResponse, httpCode int, err error,
) {
	if request.ID <= 0 {
		return response, http.StatusBadRequest, ErrNegatifID
	}

	item, err := x.Repo.GetItemByID(ctx, repo.GetItemByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		log.Printf("[GetItemByID][GetItemByID] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if item.ID == 0 {
		return response, http.StatusNotFound, ErrIDNotFound
	}

	response = GetItemByIDResponse{
		ID:   item.ID,
		Name: item.Name,
		Data: ItemDetails{
			ItemDetailsID: item.Data.ItemDetailsID,
			CPUModel:      item.Data.CPUModel,
			RAM: func() *string {
				if item.Data.RAM == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *item.Data.RAM)
				return &str
			}(),
			Year: item.Data.Year,
			ScreenSize: func() *string {
				if item.Data.ScreenSize == nil {
					return nil
				}

				str := fmt.Sprintf("%.1f inch", *item.Data.ScreenSize)
				return &str
			}(),
			Capacity: func() *string {
				if item.Data.Capacity == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *item.Data.Capacity)
				return &str
			}(),
			Color: item.Data.Color,
			Price: item.Data.Price,
		},
	}

	return response, http.StatusOK, nil
}
