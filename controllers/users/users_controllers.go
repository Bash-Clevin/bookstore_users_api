package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bash-Clevin/bookstore_users_api/domain/users"
	"github.com/Bash-Clevin/bookstore_users_api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "")
// }
