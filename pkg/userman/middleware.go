package userman

import (
	"github.com/gin-gonic/gin"
)

func TokenValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("token")
		}
		c.Set("token", token)

		// user, err := GetUserByBolorID(UserID)
		// if err != nil {
		// 	c.AbortWithError(http.StatusForbidden, err)
		// 	return
		// }
		// if user.RegAt == "" {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		// c.Set("userman", *user)
		c.Next()
	}

}
