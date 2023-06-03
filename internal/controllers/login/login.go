package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msik-404/micro-appoint-employees/internal/auth"
	"github.com/msik-404/micro-appoint-employees/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginEndPoint(db *mongo.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type LoginPlain struct {
			Mail     string `json:"mail" binding:"max=30"`
			PlainPwd string `json:"pwd" bidning:"max=72"`
			IsOwner  bool   `json:"is_owner"`
		}
		var loginPlain LoginPlain
		if err := c.BindJSON(&loginPlain); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var result *mongo.SingleResult
		if loginPlain.IsOwner {
			result = models.FindOneOwner(db, loginPlain.Mail)
		} else {
			result = models.FindOneCustomer(db, loginPlain.Mail)
		}
		type Credentials struct {
			Id        primitive.ObjectID `bson:"_id"`
			HashedPwd []byte             `bson:"pwd"`
		}
		var creds Credentials
		err := result.Decode(&creds)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithError(http.StatusUnauthorized, err)
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}
		isAuthorized, err := auth.AreEqual(creds.HashedPwd, []byte(loginPlain.PlainPwd))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if isAuthorized {
            // tokenStr, err := auth.CreateJWT(creds.Id, loginPlain.IsOwner)
            if err != nil {
			    c.AbortWithError(http.StatusUnauthorized, err)
                return
            }
            // c.SetSameSite(http.SameSiteLaxMode)
            // c.SetCookie("Authorization", tokenStr, 3600)
		} else {
			c.AbortWithError(http.StatusUnauthorized, err)
		}
	}
	return gin.HandlerFunc(fn)
}
