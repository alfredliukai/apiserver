package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"errors"
	"fmt"
)

var(
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Context struct {
	ID			uint64
	Username	string
}

func secretFunc(secret string) jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string,secret string)(*Context,error){
	ctx := &Context{}
	fmt.Println("tokenString",tokenString)
	token,err := jwt.Parse(tokenString,secretFunc(secret))

	if err!=nil{
		return ctx, err
		// Read the token if it's valid.
	}else if claims,ok := token.Claims.(jwt.MapClaims);ok && token.Valid{
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx,nil
	}else {
		return ctx,err
	}
}

func ParseRequest(ctx *gin.Context)(*Context,error){
	header := ctx.Request.Header.Get("Authorization")
	secret := viper.GetString("jwt_secret")
	if len(header) == 0{
		return &Context{},ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header,"Bearer %s", &t)
	return Parse(t,secret)
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context,c Context,secret string)(tokenString string,err error){
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret ==""{
		secret = viper.GetString("jwt_secret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":		c.ID,
		"username":	c.Username,
		"nbf":		time.Now().Unix(),
		"iat":		time.Now().Unix(),
	})
	tokenString,err = token.SignedString([]byte(secret))
	return
}