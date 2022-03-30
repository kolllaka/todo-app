package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		userCtx: userID,
	})

}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
