package handler

import (
	"net/http"

	"github.com/KoLLlaka/todo-app/internal/todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	
}

func (h *Handler) signIn(c *gin.Context) {

}
