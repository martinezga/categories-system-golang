package api

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/containers"
)

func ServeApi() {
	router := gin.Default()
	containers.BuildCategoryContainer(router)

	router.Run()
}
