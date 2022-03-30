package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// parse token
	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {
	userID, ok := c.Get(userCtx)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return userIDInt, nil
}
