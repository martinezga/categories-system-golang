package api

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/containers"
	"gorm.io/gorm"
)

func ServeApi(db *gorm.DB) {
	router := gin.Default()
	containers.BuildCategoryContainer(router, db)

	router.Run()
}
