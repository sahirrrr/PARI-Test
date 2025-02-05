package create_new_item

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
	CreateNewItemSvc interface {
		CreateNewItem(
			/*req*/ ctx context.Context, request CreateNewItemRequest) (
			/*res*/ response CreateNewItemResponse, httpCode int, err error,
		)
	}

	CreateNewItemRequest struct {
		Name       string   `json:"name"`
		CPUModel   *string  `json:"cpu_model"`
		RAM        *int64   `json:"ram"`
		Year       *int64   `json:"year"`
		ScreenSize *float64 `json:"screen_size"`
		Capacity   *int64   `json:"capacity"`
		Color      *string  `json:"color"`
		Price      *float64 `json:"price"`
	}

	CreateNewItemResponse struct {
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
	ErrItemNameEmpty          = errors.New("error item name cannot be empty")
	ErrCPUModelEmpty          = errors.New("error cpu model cannot be empty")
	ErrColorEmpty             = errors.New("error color cannot be empty")
	ErrRamLessThanZero        = errors.New("error ram cannot be less than 0")
	ErrYearLessThanZero       = errors.New("error year cannot be less than 0")
	ErrScreenSizeLessThanZero = errors.New("error screen size cannot be less than 0")
	ErrPriceLessThanZero      = errors.New("error price cannot be less than 0")
	ErrDuplicateItemName      = errors.New("error duplicate item name")
)

func New(impl Dependencies) CreateNewItemSvc {
	return &Service{Dependencies: impl, Repo: repo.NewPostgreSQL(impl.PostgreSQL)}
}

func (x *Service) CreateNewItem(
	/*req*/ ctx context.Context, request CreateNewItemRequest) (
	/*res*/ response CreateNewItemResponse, httpCode int, err error,
) {
	if request.Name == "" {
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

	if request.Color != nil && *request.Color == "" {
		return response, http.StatusBadRequest, ErrItemNameEmpty
	}

	if isNegativeOrZero(request.Price) {
		return response, http.StatusBadRequest, ErrScreenSizeLessThanZero
	}

	// Check name already exist/duplicate.
	item, err := x.Repo.IsItemNameExist(ctx, repo.IsItemNameExistRequest{
		Name: request.Name,
	})
	if err != nil {
		log.Printf("[CreateNewItem][IsItemNameExist] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if item.Exist {
		return response, http.StatusBadRequest, ErrDuplicateItemName
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

	insertItem, err := pgTXC.InsertItem(ctx, repo.InsertItemRequest{
		Name: request.Name,
	})
	if err != nil {
		log.Printf("[CreateNewItem][InsertItem] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	_, err = pgTXC.InsertItemDetails(ctx, repo.InsertItemDetailsRequest{
		ItemID:     insertItem.ID,
		CPUModel:   request.CPUModel,
		RAM:        request.RAM,
		Year:       request.Year,
		ScreenSize: request.ScreenSize,
		Capacity:   request.Capacity,
		Color:      request.Color,
		Price:      request.Price,
	})
	if err != nil {
		log.Printf("[CreateNewItem][InsertItemDetails] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[CreateNewItem][Commit] error : %v", err)

		return
	}

	res, err := x.Repo.GetItemByID(ctx, repo.GetItemByIDRequest{
		ID: insertItem.ID,
	})
	if err != nil {
		log.Printf("[CreateNewItem][GetItemByID] error : %v", err)

		return response, http.StatusInternalServerError, err
	}

	response = CreateNewItemResponse{
		ID:   res.ID,
		Name: res.Name,
		Data: ItemDetails{
			ItemDetailsID: res.Data.ItemDetailsID,
			CPUModel:      res.Data.CPUModel,
			RAM: func() *string {
				if res.Data.RAM == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *res.Data.RAM)
				return &str
			}(),
			Year: res.Data.Year,
			ScreenSize: func() *string {
				if res.Data.ScreenSize == nil {
					return nil
				}

				str := fmt.Sprintf("%.1f inch", *res.Data.ScreenSize)
				return &str
			}(),
			Capacity: func() *string {
				if res.Data.Capacity == nil {
					return nil
				}

				str := fmt.Sprintf("%d GB", *res.Data.Capacity)
				return &str
			}(),
			Color: res.Data.Color,
			Price: res.Data.Price,
		},
	}

	return response, http.StatusCreated, nil
}

func isNegativeOrZero[T int64 | float64](val *T) bool {
	return val != nil && *val <= 0
}
