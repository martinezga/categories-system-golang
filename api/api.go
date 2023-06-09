package api

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/containers"
	"github.com/martinezga/categories-system-golang/config"
	_ "github.com/martinezga/categories-system-golang/docs/swagger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// HealthHandler
// @Summary      Check if service is running
// @Description  Health check
// @Tags         default
// @Produce      plain
// @Success      200 "up and running"
// @Router       /health [get]
func HealthHandler(c *gin.Context) {
	c.String(200, "up and running")
}

// @title           Swagger Category and Tags system API
// @version         1.0
// @description     This is a category and tags system server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func ServeApi(config config.Environment, db *gorm.DB) {
	router := gin.Default()

	if config.GinMode == "debug" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.GET("/health", HealthHandler)
	containers.BuildCategoryContainer(router, db)

	router.Run()
}
