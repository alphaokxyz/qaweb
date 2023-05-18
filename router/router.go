package router

import (
	"crypto/tls"
	"net/http"
	"qaweb/api"
	"qaweb/config"
	"qaweb/middleware"

	"github.com/gin-gonic/gin"
)

func Initrouter() {
	gin.SetMode(config.Appmode)
	r := gin.New()
	r.Use(middleware.Logmiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	router := r.Group("api")
	{
		router.GET("users", api.Getusers)
		router.POST("login", api.Login)
		router.POST("user/add", api.Adduser)
		router.GET("questions", api.Getquestions)
		router.GET("question/info/:id", api.Getquestioninfo)
		router.GET("answers/:id", api.Getanswersofaquestion)
		router.GET("answer/info/:id", api.Getanswerinfo)
	}
	jwtrouter := r.Group("api")
	jwtrouter.Use(middleware.Jwtmiddleware())
	{
		jwtrouter.PUT("user/:id", api.Edituser)
		jwtrouter.DELETE("user/:id", api.Deleteuser)
		jwtrouter.POST("question/add", api.Addquestion)
		jwtrouter.PUT("question/:id", api.Editquestion)
		jwtrouter.DELETE("question/:id", api.Deletequestion)
		jwtrouter.POST("answer/add", api.Addanswer)
		jwtrouter.PUT("answer/:id", api.Editanswer)
		jwtrouter.DELETE("answer/:id", api.Deleteanswer)
	}

	server := &http.Server{
		Addr:    config.Httpport,
		Handler: r,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{},
		},
	}

	server.ListenAndServeTLS("/etc/letsencrypt/live/api.alphaokxyz.site/fullchain.pem", "/etc/letsencrypt/live/api.alphaokxyz.site/privkey.pem")
}
