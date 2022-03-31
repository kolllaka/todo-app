package handler

import (
	"net/http"
	"strconv"

	"github.com/KoLLlaka/todo-app/internal/todo"

	"github.com/gin-gonic/gin"
)

const (
	paramID = "id"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	listID, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id": listID,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	list, err := h.services.TodoList.GetByID(userID, listID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list": list,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	var updateInput todo.UpdateListInput
	if err := c.BindJSON(&updateInput); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userID, listID, updateInput); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		"ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param(paramID))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid listID param")
		return
	}

	if err := h.services.TodoList.Delete(userID, listID); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
