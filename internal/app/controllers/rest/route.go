package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahirrrr/PARI-Test/internal/app/service/create_new_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/delete_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_item_by_id"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_list_items"
	"github.com/sahirrrr/PARI-Test/internal/app/service/update_item"
	"go.uber.org/dig"
)

type (
	ServicesImpl struct {
		Services
	}

	Services struct {
		dig.In
		get_list_items.GetListItemsSvc
		get_item_by_id.GetItemByIDSvc
		create_new_item.CreateNewItemSvc
		delete_item.DeleteItemSvc
		update_item.UpdateItemSvc
	}
)

func SetRoute(
	e *echo.Echo,
	services Services,
) {
	handler := ServicesImpl{Services: services}

	// Health
	e.GET("/health", healthCheck)

	// Items
	e.GET("/items", handler.getListItems)
	e.GET("/item", handler.getItemByID)
	e.POST("/item", handler.createNewItem)
	e.DELETE("/item", handler.deleteItem)
	e.PUT("/item", handler.updateItem)

}

func healthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, "OK")
}
