package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func AreEqual(hashedPwd []byte, plainPwd []byte) (bool, error) {
	byteHash := hashedPwd
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateJWT(
    userID primitive.ObjectID, 
    isOwner bool,
) (string, error) {
    token := jwt.New(jwt.SigningMethodES256)

    claims := token.Claims.(jwt.MapClaims)
    claims["exp"] = time.Now().Add(time.Hour).Unix()
    claims["user_id"] = userID
    claims["is_owner"] = isOwner

    tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
        return "", err
    }
    return tokenStr, err
}
