package handler

import (
	"net/http"
	"strconv"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	itemID, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"item_id": itemID,
	})
}

func (h *Handler) getAllItem(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
