package middleware

import (
	"net/http"
	"qaweb/config"
	"qaweb/errormessage"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Jwtkey = []byte(config.Jwtkey)

type Requestin struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Maketoken(username string) (string, int) {
	lasttime := time.Now().Add(7200 * time.Hour)
	Makerequestin := Requestin{
		username,
		jwt.StandardClaims{
			ExpiresAt: lasttime.Unix(),
			Issuer:    "qaweb",
		},
	}
	requestinjwt := jwt.NewWithClaims(jwt.SigningMethodHS256, Makerequestin)
	token, err := requestinjwt.SignedString(Jwtkey)
	if err != nil {
		return "", errormessage.ERROR
	}
	return token, errormessage.SUCCESS
}
func Checktoken(token string) (*Requestin, int) {
	var request Requestin
	maketoken, err := jwt.ParseWithClaims(token, &request, func(token *jwt.Token) (i interface{}, e error) {
		return Jwtkey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errormessage.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errormessage.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errormessage.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if maketoken != nil {
		if key, ok := maketoken.Claims.(*Requestin); ok && maketoken.Valid {
			return key, errormessage.SUCCESS
		} else {
			return nil, errormessage.ERROR_TOKEN_WRONG
		}
	}
	return nil, errormessage.ERROR_TOKEN_WRONG
}

func Jwtmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenheader := c.Request.Header.Get("Authorization")
		if tokenheader == "" {
			code = errormessage.ERROR_TOKEN_NOTEXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormessage.Geterrormessage(code),
			})
			c.Abort()
			return
		}
		checktoken := strings.Split(tokenheader, " ")
		if len(checktoken) == 0 {
			code = errormessage.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormessage.Geterrormessage(code),
			})
			c.Abort()
			return
		}

		if len(checktoken) != 2 && checktoken[0] != "Bearer" {
			code = errormessage.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormessage.Geterrormessage(code),
			})
			c.Abort()
			return
		}
		key, newcode := Checktoken(checktoken[1])
		if newcode != errormessage.SUCCESS {
			code = newcode
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errormessage.Geterrormessage(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key)
		c.Next()
	}
}

func Getusername(token string) string {
	var request Requestin
	maketoken, err := jwt.ParseWithClaims(token, &request, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return ""
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return ""
			} else {
				return ""
			}
		}
	}
	if maketoken != nil {
		if key, ok := maketoken.Claims.(*Requestin); ok && maketoken.Valid {
			return key.Username
		} else {
			return ""
		}
	}
	return ""
}
