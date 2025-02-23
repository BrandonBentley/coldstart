package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Check(c *gin.Context) {

	resp, err := h.healthService.Check(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
