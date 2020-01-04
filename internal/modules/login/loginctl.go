package login

import (
	"net/http"

	"github.com/Raybird/whale/internal/auth"
	"github.com/Raybird/whale/internal/models"
	"github.com/Raybird/whale/internal/modules"
	"github.com/Raybird/whale/internal/responses"
	"github.com/Raybird/whale/internal/utils/formaterror"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Ctrl ...
type Ctrl struct {
	Base *modules.BaseCtrl
}

// Login ...
func (ctl *Ctrl) Login(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := ctl.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(c.Writer, http.StatusUnprocessableEntity, formattedError)
		return
	}
	c.JSON(http.StatusOK, token)
}

// SignIn ...
func (ctl *Ctrl) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = ctl.Base.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
