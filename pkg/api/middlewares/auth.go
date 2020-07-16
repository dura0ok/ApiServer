package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func RequireAuthentication() func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(h http.ResponseWriter, req *http.Request) {
			authorization := req.Header.Get("Authorization")
			authorizationSlice := strings.Split(authorization, " ")
			if len(authorizationSlice) < 2 || authorizationSlice[0] != "Bearer" || authorizationSlice[1] == ""{
				http.Error(h, "Not authorized", 401)
				return
			}
			token := authorizationSlice[1]
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil{
				http.Error(h, "Not aaa", 401)
				return
			}
			delete(claims, "exp")
			ctx := context.WithValue(req.Context(), "user", claims)
			fmt.Println("Creating context :D")
			next.ServeHTTP(h, req.WithContext(ctx))
		})
	}
}
