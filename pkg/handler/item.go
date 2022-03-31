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
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	itemID, err := strconv.Atoi(c.Param(paramItemID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid itemID param")
		return
	}

	item, err := h.services.TodoItem.GetByID(userID, listID, itemID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"item": item,
	})
}

func (h *Handler) updateItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	itemID, err := strconv.Atoi(c.Param(paramItemID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid itemID param")
		return
	}

	var updateInput todo.UpdateItemInput
	if err := c.BindJSON(&updateInput); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userID, listID, itemID, updateInput); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		"ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	itemID, err := strconv.Atoi(c.Param(paramItemID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid itemID param")
		return
	}

	if err := h.services.TodoItem.Delete(userID, listID, itemID); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
