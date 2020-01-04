package users

import (
	"net/http"

	"github.com/Raybird/whale/internal/models"
	"github.com/Raybird/whale/internal/modules"
	"github.com/Raybird/whale/internal/responses"
	"github.com/Raybird/whale/internal/utils/formaterror"
	"github.com/gin-gonic/gin"
)

// Ctrl ...
type Ctrl struct {
	Base *modules.BaseCtrl
}

// CreateUser ...
func (ctl *Ctrl) CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(ctl.Base.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(c.Writer, http.StatusInternalServerError, formattedError)
		return
	}

	c.JSON(http.StatusOK, userCreated)
}

// GetUsers ...
func (ctl *Ctrl) GetUsers(c *gin.Context) {
	user := models.User{}

	users, err := user.FindAllUsers(ctl.Base.DB)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
