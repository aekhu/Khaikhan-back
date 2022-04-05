package common

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ServerError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	App.ErrorLog.Println(trace)
	c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func ClientError(c *gin.Context, code int, err error) {
	c.String(code, err.Error())
}
