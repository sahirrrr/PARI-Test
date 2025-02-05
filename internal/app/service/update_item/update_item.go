package update_item

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
	UpdateItemSvc interface {
		UpdateItem(
			/*req*/ ctx context.Context, request UpdateItemRequest) (
			/*res*/ response UpdateItemResponse, httpCode int, err error,
		)
	}

	UpdateItemRequest struct {
		ID         int64    `json:"id"`
		Name       *string  `json:"name"`
		CPUModel   *string  `json:"cpu_model"`
		RAM        *int64   `json:"ram"`
		Year       *int64   `json:"year"`
		ScreenSize *float64 `json:"screen_size"`
		Capacity   *int64   `json:"capacity"`
		Color      *string  `json:"color"`
		Price      *float64 `json:"price"`
	}

	UpdateItemResponse struct {
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
	ErrNegatifID              = errors.New("error id cannot be negatif")
	ErrIDNotFound             = errors.New("error id cannot be found")
	ErrItemNameEmpty          = errors.New("error item name cannot be empty")
	ErrCPUModelEmpty          = errors.New("error cpu model cannot be empty")
	ErrRamLessThanZero        = errors.New("error ram cannot be less than 0")
	ErrYearLessThanZero       = errors.New("error year cannot be less than 0")
	ErrScreenSizeLessThanZero = errors.New("error screen size cannot be less than 0")
	ErrCapacityLessThanZero   = errors.New("error capacity cannot be less than 0")
	ErrColorEmpty             = errors.New("error color cannot be empty")
	ErrPriceLessThanZero      = errors.New("error price cannot be less than 0")
	ErrDuplicateItemName      = errors.New("error duplicate item name")
)

func New(impl Dependencies) UpdateItemSvc {
	return &Service{Dependencies: impl, Repo: repo.NewPostgreSQL(impl.PostgreSQL)}
}

func (x *Service) UpdateItem(
	/*req*/ ctx context.Context, request UpdateItemRequest) (
	/*res*/ response UpdateItemResponse, httpCode int, err error,
) {
	if request.ID <= 0 {
		return response, http.StatusBadRequest, ErrNegatifID
	}

	if request.Name != nil && *request.Name == "" {
		return response, http.StatusBadRequest, ErrItemNameEmpty
	}

	if request.CPUModel != nil && *request.CPUModel == "" {
		return response, http.StatusBadRequest, ErrCPUModelEmpty
	}

	if isNegativeOrZero(request.RAM) {
		return response, http.StatusBadRequest, ErrRamLessThanZero
	}

	if isNegativeOrZero(request.Year) {
		return response, http.StatusBadRequest, ErrYearLessThanZero
	}

	if isNegativeOrZero(request.ScreenSize) {
		return response, http.StatusBadRequest, ErrScreenSizeLessThanZero
	}

	if isNegativeOrZero(request.Capacity) {
		return response, http.StatusBadRequest, ErrScreenSizeLessThanZero
	}

	if request.Color != nil && *request.Color == "" {
		return response, http.StatusBadRequest, ErrColorEmpty
	}

	if isNegativeOrZero(request.Price) {
		return response, http.StatusBadRequest, ErrScreenSizeLessThanZero
	}

	item, err := x.Repo.GetItemByID(ctx, repo.GetItemByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		log.Printf("[UpdateItem][GetItemByID] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if item.ID == 0 {
		return response, http.StatusNotFound, ErrIDNotFound
	}

	var tx *sqlx.Tx

	tx, err = x.PostgreSQL.BeginTxx(ctx, nil)
	if err != nil {
		log.Printf("[CreateNewItem][BeginTxx] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Printf("[CreateNewItem][Rollback] error : %v", err)

				return
			}

			return
		}
	}()

	txc := repo.SQLTxConn(tx)
	pgTXC := repo.NewPostgreSQL(txc)

	// New Name
	if request.Name != nil {
		if item.Name != *request.Name {
			// Check if name already exist
			item, err := x.Repo.IsItemNameExist(ctx, repo.IsItemNameExistRequest{
				Name: *request.Name,
			})
			if err != nil {
				log.Printf("[UpdateItem][IsItemNameExist] error : %v", err)

				return response, http.StatusInternalServerError, err
			}

			if item.Exist {
				return response, http.StatusBadRequest, ErrDuplicateItemName
			}

			// update name
			_, err = pgTXC.UpdateItem(ctx, repo.UpdateItemRequest{
				ID:   request.ID,
				Name: *request.Name,
			})
			if err != nil {
				log.Printf("[UpdateItem][UpdateItem] error : %v", err)

				return response, http.StatusInternalServerError, err
			}
		}
	}

	// update details
	_, err = pgTXC.UpdateItemDetails(ctx, repo.UpdateItemDetailsRequest{
		ID:         item.Data.ItemDetailsID,
		CPUModel:   request.CPUModel,
		RAM:        request.RAM,
		Year:       request.Year,
		ScreenSize: request.ScreenSize,
		Capacity:   request.Capacity,
		Color:      request.Color,
		Price:      request.Price,
	})
	if err != nil {
		log.Printf("[UpdateItem][UpdateItemDetails] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[CreateNewItem][Commit] error : %v", err)

		return
	}

	itemUpdated, err := x.Repo.GetItemByID(ctx, repo.GetItemByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		log.Printf("[UpdateItem][GetItemByID] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	response = UpdateItemResponse{
		ID:   itemUpdated.ID,
		Name: itemUpdated.Name,
		Data: ItemDetails{
			ItemDetailsID: itemUpdated.Data.ItemDetailsID,
			CPUModel:      itemUpdated.Data.CPUModel,
			RAM: func() *string {
				if itemUpdated.Data.RAM == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *itemUpdated.Data.RAM)
				return &str
			}(),
			Year: itemUpdated.Data.Year,
			ScreenSize: func() *string {
				if itemUpdated.Data.ScreenSize == nil {
					return nil
				}

				str := fmt.Sprintf("%.1f inch", *itemUpdated.Data.ScreenSize)
				return &str
			}(),
			Capacity: func() *string {
				if itemUpdated.Data.Capacity == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *itemUpdated.Data.Capacity)
				return &str
			}(),
			Color: itemUpdated.Data.Color,
			Price: itemUpdated.Data.Price,
		},
	}

	return response, http.StatusOK, nil
}

func isNegativeOrZero[T int64 | float64](val *T) bool {
	return val != nil && *val <= 0
}
