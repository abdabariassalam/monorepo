package http

import (
	"net/http"

	"github.com/bariasabda/monorepo/packages/fetch/constants"
	"github.com/bariasabda/monorepo/packages/fetch/domain/service"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	svc service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *httpHandler {
	return &httpHandler{
		svc: service,
	}
}

func (h *httpHandler) VerifyToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Access-Token")
	if token == "" {
		c.JSON(http.StatusForbidden, constants.ErrTokenRequired)
		return
	}
	user, err := h.svc.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, constants.ErrBadToken)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *httpHandler) Aggregator(c *gin.Context) {
	token := c.Request.Header.Get("X-Access-Token")
	if token == "" {
		c.JSON(http.StatusForbidden, constants.ErrTokenRequired)
		return
	}
	resp, err := h.svc.Aggregator(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *httpHandler) CurrencyConverter(c *gin.Context) {
	token := c.Request.Header.Get("X-Access-Token")
	if token == "" {
		c.JSON(http.StatusForbidden, constants.ErrTokenRequired)
		return
	}
	resp, err := h.svc.CurrencyConverter(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
