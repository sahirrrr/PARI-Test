package get_list_items

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
	GetListItemsSvc interface {
		GetListItems(
			/*req*/ ctx context.Context, request GetListItemsRequest) (
			/*res*/ response []GetListItemsResponse, httpCode int, err error,
		)
	}

	GetListItemsRequest struct {
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}

	GetListItemsResponse struct {
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
	ErrNegatifLimit  = errors.New("error limit cannot be negatif")
	ErrNegatifOffSet = errors.New("error offset cannot be negatif")
	ErrMinimumLimit  = errors.New("error minimum limit 10")
	ErrMaxsimumLimit = errors.New("error maxsimum limit 100")
)

func New(impl Dependencies) GetListItemsSvc {
	return &Service{Dependencies: impl, Repo: repo.NewPostgreSQL(impl.PostgreSQL)}
}

func (x *Service) GetListItems(
	/*req*/ ctx context.Context, request GetListItemsRequest) (
	/*res*/ response []GetListItemsResponse, httpCode int, err error,
) {
	if request.Limit < 0 {
		return response, http.StatusBadRequest, ErrNegatifLimit
	}

	if request.Offset < 0 {
		return response, http.StatusBadRequest, ErrNegatifOffSet
	}

	if request.Limit >= 0 && request.Limit < 10 {
		return response, http.StatusBadRequest, ErrMinimumLimit
	}

	if request.Limit > 100 {
		return response, http.StatusBadRequest, ErrMaxsimumLimit
	}

	items, err := x.Repo.GetListItems(ctx, repo.GetListItemsRequest{
		Limit:  request.Limit,
		Offset: request.Offset,
	})
	if err != nil {
		log.Printf("[GetListItems][GetListItems] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	for _, item := range items {
		list := GetListItemsResponse{
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
		response = append(response, list)
	}

	return response, http.StatusOK, nil
}
