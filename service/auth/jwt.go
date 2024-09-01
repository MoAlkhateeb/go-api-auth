package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MoAlkhateeb/go-api-auth/config"
	"github.com/MoAlkhateeb/go-api-auth/types"
	"github.com/MoAlkhateeb/go-api-auth/utils"
	"github.com/golang-jwt/jwt/v5"
)

const userKey string = "userID"

func CreateJWT(secret []byte, userID uint) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.FormatUint(uint64(userID), 10),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)

		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("Invalid validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("Invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims[userKey].(string)

		userID, _ := strconv.Atoi(str)
		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("Failed to get user by ID: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("Permission Denied"))
}

func GetIDFromContext(ctx context.Context) int {
	userID := ctx.Value(userKey)
	if userID == nil {
		return -1
	}
	return userID.(int)
}
