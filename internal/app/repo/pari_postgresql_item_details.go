package repo

import (
	"context"

	"github.com/sahirrrr/PARI-Test/internal/app/repo/postgresql_query"
	"github.com/sahirrrr/PARI-Test/pkg/entity"
)

type (
	ItemDetailsRespository interface {
		InsertItemDetails(
			/*req*/ ctx context.Context, request InsertItemDetailsRequest) (
			/*res*/ response InsertItemDetailsResponse, err error,
		)
		UpdateItemDetails(
			/*req*/ ctx context.Context, request UpdateItemDetailsRequest) (
			/*res*/ response UpdateItemDetailsResponse, err error,
		)
	}

	InsertItemDetailsRequest struct {
		ItemID     int64
		CPUModel   *string
		RAM        *int64
		Year       *int64
		ScreenSize *float64
		Capacity   *int64
		Color      *string
		Price      *float64
	}

	InsertItemDetailsResponse struct {
		ID int64
	}

	UpdateItemDetailsRequest struct {
		ID         int64
		CPUModel   *string
		RAM        *int64
		Year       *int64
		ScreenSize *float64
		Capacity   *int64
		Color      *string
		Price      *float64
	}

	UpdateItemDetailsResponse struct {
		RowsAffected int
	}
)

func (x *postgresql) InsertItemDetails(
	/*req*/ ctx context.Context, request InsertItemDetailsRequest) (
	/*res*/ response InsertItemDetailsResponse, err error,
) {
	query := postgresql_query.InsertItemDetails

	args := entity.List{
		request.ItemID,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	if request.CPUModel != nil {
		args[1] = request.CPUModel
	}

	if request.RAM != nil {
		args[2] = request.RAM
	}

	if request.Year != nil {
		args[3] = request.Year
	}

	if request.ScreenSize != nil {
		args[4] = request.ScreenSize
	}

	if request.Capacity != nil {
		args[5] = request.Capacity
	}

	if request.Color != nil {
		args[6] = request.Color
	}

	if request.Price != nil {
		args[7] = request.Price
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

func (x *postgresql) UpdateItemDetails(
	/*req*/ ctx context.Context, request UpdateItemDetailsRequest) (
	/*res*/ response UpdateItemDetailsResponse, err error,
) {
	query := postgresql_query.UpdateItemDetails

	args := entity.List{
		request.ID,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	if request.CPUModel != nil {
		args[1] = request.CPUModel
	}

	if request.RAM != nil {
		args[2] = request.RAM
	}

	if request.Year != nil {
		args[3] = request.Year
	}

	if request.ScreenSize != nil {
		args[4] = request.ScreenSize
	}

	if request.Capacity != nil {
		args[5] = request.Capacity
	}

	if request.Color != nil {
		args[6] = request.Color
	}

	if request.Price != nil {
		args[7] = request.Price
	}

	if err = new(SQL).BoxExec(x.tc.ExecContext(ctx, query, args...)).Scan(&response.RowsAffected, nil); err != nil {
		err = new(entity.SourceError).With(err, ctx, request)
	}

	return response, err
}
