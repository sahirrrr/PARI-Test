package rest

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sahirrrr/PARI-Test/internal/app/service/create_new_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/delete_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_item_by_id"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_list_items"
	"github.com/sahirrrr/PARI-Test/internal/app/service/update_item"
)

var (
	ErrInvalidLimit  = errors.New("invalid limit value")
	ErrInvalidOffset = errors.New("invalid offset value")
)

func (x *ServicesImpl) getListItems(c echo.Context) (err error) {
	ctx := c.Request().Context()
	limit := c.QueryParam("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("Invalid offset: %v", err)

		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidLimit)))
	}

	offset := c.QueryParam("offset")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("Invalid offset: %v", err)

		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidOffset)))
	}

	request := get_list_items.GetListItemsRequest{
		Limit:  int64(limitInt),
		Offset: int64(offsetInt),
	}

	response, httpCode, err := x.GetListItemsSvc.GetListItems(ctx, request)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: response})
}

func (x *ServicesImpl) getItemByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.QueryParam("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid id: %v", err)

		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidLimit)))
	}

	request := get_item_by_id.GetItemByIDRequest{
		ID: int64(idInt),
	}

	response, httpCode, err := x.GetItemByIDSvc.GetItemByID(ctx, request)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: response})
}

func (x *ServicesImpl) createNewItem(c echo.Context) (err error) {
	ctx := c.Request().Context()

	request := create_new_item.CreateNewItemRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
	}

	response, httpCode, err := x.CreateNewItemSvc.CreateNewItem(ctx, request)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: response})
}

func (x *ServicesImpl) deleteItem(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.QueryParam("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid id: %v", err)

		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidLimit)))
	}

	request := delete_item.DeleteItemRequest{
		ID: int64(idInt),
	}

	response, httpCode, err := x.DeleteItemSvc.DeleteItem(ctx, request)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: response})
}

func (x *ServicesImpl) updateItem(c echo.Context) (err error) {
	ctx := c.Request().Context()

	request := update_item.UpdateItemRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
	}

	response, httpCode, err := x.UpdateItemSvc.UpdateItem(ctx, request)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: response})
}
