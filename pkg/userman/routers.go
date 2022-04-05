package userman

import "github.com/gin-gonic/gin"

func Routers(route *gin.Engine) {
	// route.GET("/login", LoginURLHandler)
	route.GET("/authenticate", AuthenticateHandler)
	route.POST("/authenticate", AuthenticateHandler)
	route.POST("/checkToken", TokenValidate(), CheckTokenHandler)
	route.POST("/logout", TokenValidate(), LogoutHandler)
}
