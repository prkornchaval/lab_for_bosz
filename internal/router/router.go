package router

import (
	"labForBosz/internal/handler"
	"labForBosz/property"
	"net/http"

	"github.com/centraldigital/cfw-core-lib/pkg/util/ginutil"
	"github.com/gin-gonic/gin"

	docs "labForBosz/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const serviceBaseURL = "/customer-api"

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
// @description API key
func Setup(r *gin.Engine, hdl handler.Handler) {
	docs.SwaggerInfo.Title = property.Get().Service.ServiceName
	docs.SwaggerInfo.Description = property.Get().Service.ServiceDescription
	docs.SwaggerInfo.Version = property.Get().Service.ApiDocsVersion
	docs.SwaggerInfo.Host = property.Get().Service.ApiDocsHost
	docs.SwaggerInfo.Schemes = []string{property.Get().Service.ApiDocsSchema}
	docs.SwaggerInfo.BasePath = serviceBaseURL

	basePath := r.Group(serviceBaseURL)

	v1 := basePath.Group("/v1")
	v1.POST("/", func(c *gin.Context) { ginutil.BindReqJson200Resp(c, hdl.CreateCustomer) })
	v1.POST("customer-address", func(c *gin.Context) { ginutil.BindReqJson200Resp(c, hdl.CreateCustomerAddress) })
	v1.GET("/:id", func(c *gin.Context) { ginutil.BindReqJson200Resp(c, hdl.GetCustomerbyId) })

	if property.Get().Service.ApiDocs {
		docsPath := "/docs" + serviceBaseURL
		r.GET(docsPath, func(ctx *gin.Context) {
			ctx.Redirect(http.StatusTemporaryRedirect, docsPath+"/swagger/index.html")
		})
		r.GET(docsPath+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
