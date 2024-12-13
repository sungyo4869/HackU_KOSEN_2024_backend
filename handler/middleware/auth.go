package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("SECRET_KEY")

type UserIDKey struct{}

func Auth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		log.Println("login ok")

		// tokenの認証
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// ペイロードの読み出し
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"].(float64)
		ctx := context.WithValue(r.Context(), UserIDKey{}, int64(userID))
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
