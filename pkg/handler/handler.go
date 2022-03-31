package handler

import (
	"github.com/KoLLlaka/todo-app/pkg/service"

	"github.com/gin-gonic/gin"
)

const (
	paramID     = "id"
	paramItemID = "item_id"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:"+paramID, h.getListById)
			lists.PUT("/:"+paramID, h.updateList)
			lists.DELETE("/:"+paramID, h.deleteList)

			items := lists.Group(":" + paramID + "/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:"+paramItemID, h.getItemById)
				items.PUT("/:"+paramItemID, h.updateItem)
				items.DELETE("/:"+paramItemID, h.deleteItem)
			}
		}
	}

	return router
}
