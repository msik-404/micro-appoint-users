package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/auth"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func AddCompanyEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type CustomerPlain struct {
			Mail     string `json:"mail" binding:"max=30"`
			PlainPwd string `json:"pwd" bidning:"max=72"`
			Name     string `json:"name,omitempty" binding:"max=30"`
			Surname  string `json:"surname,omitempty" binding:"max=30"`
		}
		var newCustomerPlain CustomerPlain
		if err := c.BindJSON(&newCustomerPlain); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		HashedPwd, err := auth.HashAndSalt([]byte(newCustomerPlain.PlainPwd))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newCustomer := models.Customer{
			Mail:      newCustomerPlain.Mail,
			HashedPwd: HashedPwd,
			Name:      newCustomerPlain.Name,
			Surname:   newCustomerPlain.Surname,
		}
		result, err := newCustomer.InsertOne(db)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
