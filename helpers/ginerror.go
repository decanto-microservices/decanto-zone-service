package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGenericError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err)
}
