package libraries

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTLib struct {
	secretKey string
	issuer    string
}

func NewJWTLib() *JWTLib {
	return &JWTLib{
		secretKey: os.Getenv("APP_ACCESS_SECRET_KEY"),
		issuer:    "issuer",
	}
}

func (lib *JWTLib) GenerateToken(payload map[string]interface{}) (string, error) {

	claims := jwt.MapClaims{}

	for k, v := range payload {
		claims[k] = v
	}

	claims["issuer"] = lib.issuer
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["iss"] = time.Now().Unix()

	wc := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := wc.SignedString([]byte(lib.secretKey))
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (lib *JWTLib) ValidateToken(token string) (*jwt.Token, error) {

	fmt.Println(token)

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte([]byte(os.Getenv("APP_ACCESS_SECRET_KEY"))), nil
	})
}
