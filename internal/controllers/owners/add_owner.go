package owners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/msik-404/micro-appoint-employees/internal/auth"
	"github.com/msik-404/micro-appoint-employees/internal/models"
)

func AddOwnerEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type OwnerPlain struct {
			Mail     string `json:"mail" binding:"max=30"`
			PlainPwd string `json:"pwd" bidning:"max=72"`
			Name     string `json:"name,omitempty" binding:"max=30"`
			Surname  string `json:"surname,omitempty" binding:"max=30"`
		}
		var newOwnerPlain OwnerPlain
		if err := c.BindJSON(&newOwnerPlain); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		hashedPwd, err := auth.HashAndSalt([]byte(newOwnerPlain.PlainPwd))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newOwner := models.Owner{
			Mail:      newOwnerPlain.Mail,
			HashedPwd: hashedPwd,
			Name:      newOwnerPlain.Name,
			Surname:   newOwnerPlain.Surname,
		}
		result, err := newOwner.InsertOne(db)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
	return gin.HandlerFunc(fn)
}
