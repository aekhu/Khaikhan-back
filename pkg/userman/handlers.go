package userman

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateHandler(c *gin.Context) {
	if c.Request.Method == http.MethodPost {

	} else if c.Request.Method == http.MethodGet {

		// _, exist := UserIDExist(buser.ID)
		// if !exist {
		// 	_, err := CreateUser(buser.ID, buser.Name, buser.Email, buser.Phone)
		// 	if err != nil {
		// 		common.ServerError(c, err)
		// 		return
		// 	}
		// }
		// c.Redirect(http.StatusSeeOther, common.App.WebURL+"?t="+accessToken)
	} else {
		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func CheckTokenHandler(c *gin.Context) {
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func LogoutHandler(c *gin.Context) {
	//token := c.MustGet("token").(string)
	//expire token
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
