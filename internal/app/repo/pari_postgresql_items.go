package repo

import (
	"context"

	"github.com/sahirrrr/PARI-Test/internal/app/repo/postgresql_query"
	"github.com/sahirrrr/PARI-Test/pkg/entity"
)

type (
	ItemsRespository interface {
		GetListItems(
			/*req*/ ctx context.Context, request GetListItemsRequest) (
			/*res*/ response []GetListItemsResponse, err error,
		)
		GetItemByID(
			/*req*/ ctx context.Context, request GetItemByIDRequest) (
			/*res*/ response GetItemByIDResponse, err error,
		)
		IsItemNameExist(
			/*req*/ ctx context.Context, request IsItemNameExistRequest) (
			/*res*/ response IsItemNameExistResponse, err error,
		)
		InsertItem(
			/*req*/ ctx context.Context, request InsertItemRequest) (
			/*res*/ response InsertItemResponse, err error,
		)
		DeleteItem(
			/*req*/ ctx context.Context, request DeleteItemRequest) (
			/*res*/ response DeleteItemResponse, err error,
		)
		UpdateItem(
			/*req*/ ctx context.Context, request UpdateItemRequest) (
			/*res*/ response UpdateItemResponse, err error,
		)
	}

	GetListItemsRequest struct {
		Limit  int64
		Offset int64
	}

	GetListItemsResponse struct {
		ID   int64
		Name string
		Data ItemDetails
	}

	ItemDetails struct {
		ItemDetailsID int64
		CPUModel      *string
		RAM           *int64
		Year          *int64
		ScreenSize    *float64
		Capacity      *int64
		Color         *string
		Price         *float64
	}

	GetItemByIDRequest struct {
		ID int64
	}

	GetItemByIDResponse struct {
		ID   int64
		Name string
		Data ItemDetails
	}

	IsItemNameExistRequest struct {
		Name string
	}

	IsItemNameExistResponse struct {
		Exist bool
	}

	InsertItemRequest struct {
		Name string
	}

	InsertItemResponse struct {
		ID int64
	}

	DeleteItemRequest struct {
		ID int64
	}

	DeleteItemResponse struct {
		RowsAffected int
	}

	UpdateItemRequest struct {
		ID   int64
		Name string
	}

	UpdateItemResponse struct {
		RowsAffected int
	}
)

func (x *postgresql) GetListItems(
	/*req*/ ctx context.Context, request GetListItemsRequest) (
	/*res*/ response []GetListItemsResponse, err error,
) {
	query := postgresql_query.GetListItems

	args := entity.List{
		request.Limit,
		request.Offset,
	}

	row := func(i int) entity.List {
		response = append(response, GetListItemsResponse{})

		return entity.List{
			&response[i].ID,
			&response[i].Name,
			&response[i].Data.ItemDetailsID,
			&response[i].Data.CPUModel,
			&response[i].Data.RAM,
			&response[i].Data.Year,
			&response[i].Data.ScreenSize,
			&response[i].Data.Capacity,
			&response[i].Data.Color,
			&response[i].Data.Price,
		}
	}

	if err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}

func (x *postgresql) GetItemByID(
	/*req*/ ctx context.Context, request GetItemByIDRequest) (
	/*res*/ response GetItemByIDResponse, err error,
) {
	query := postgresql_query.GetItemByID

	args := entity.List{
		request.ID,
	}

	row := func(i int) entity.List {
		if i > 0 {
			return nil
		}

		return entity.List{
			&response.ID,
			&response.Name,
			&response.Data.ItemDetailsID,
			&response.Data.CPUModel,
			&response.Data.RAM,
			&response.Data.Year,
			&response.Data.ScreenSize,
			&response.Data.Capacity,
			&response.Data.Color,
			&response.Data.Price,
		}
	}

	if err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}

func (x *postgresql) IsItemNameExist(
	/*req*/ ctx context.Context, request IsItemNameExistRequest) (
	/*res*/ response IsItemNameExistResponse, err error,
) {
	query := postgresql_query.IsItemNameExist

	args := entity.List{
		request.Name,
	}

	row := func(i int) entity.List {
		if i > 0 {
			return nil
		}

		return entity.List{
			&response.Exist,
		}
	}

	if err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}

func (x *postgresql) InsertItem(
	/*req*/ ctx context.Context, request InsertItemRequest) (
	/*res*/ response InsertItemResponse, err error,
) {
	query := postgresql_query.InsertItem

	args := entity.List{
		request.Name,
	}

	row := func(i int) entity.List {
		if i > 0 {
			return nil
		}

		return entity.List{
			&response.ID,
		}
	}

	if err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}

func (x *postgresql) DeleteItem(
	/*req*/ ctx context.Context, request DeleteItemRequest) (
	/*res*/ response DeleteItemResponse, err error,
) {
	query := postgresql_query.DeleteItem

	args := entity.List{
		request.ID,
	}

	if err = new(SQL).BoxExec(x.tc.ExecContext(ctx, query, args...)).Scan(&response.RowsAffected, nil); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}

func (x *postgresql) UpdateItem(
	/*req*/ ctx context.Context, request UpdateItemRequest) (
	/*res*/ response UpdateItemResponse, err error,
) {
	query := postgresql_query.UpdateItem

	args := entity.List{
		request.ID,
		request.Name,
	}

	if err = new(SQL).BoxExec(x.tc.ExecContext(ctx, query, args...)).Scan(&response.RowsAffected, nil); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}
