package http

import "github.com/gin-gonic/gin"

func Config(r *gin.Engine) {
	notLogin := r.Group("/api/ams-ce")
	{
		notLogin.GET("/ping", ping)
	}

	userLogin := r.Group("/api/ams-ce").Use(shouldBeLogin())
	{
		userLogin.GET("/hosts", hostGets)
		userLogin.POST("/hosts", hostPost)
		userLogin.GET("/host/:id", hostGet)
		userLogin.PUT("/hosts/tenant", hostTenantPut)
		userLogin.PUT("/hosts/back", hostBackPut)
		userLogin.PUT("/hosts/note", hostNotePut)
		userLogin.PUT("/hosts/cate", hostCatePut)
		userLogin.DELETE("/hosts", hostDel)
		userLogin.GET("/hosts/search", hostSearchGets)
	}

	v1 := r.Group("/v1/ams-ce").Use(shouldBeService())
	{
		v1.POST("/hosts/register", v1HostRegister)
	}
}
