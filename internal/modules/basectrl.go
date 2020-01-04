package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// BaseCtrl ...
type BaseCtrl struct {
	DB     *gorm.DB
	Router *gin.Engine
}
