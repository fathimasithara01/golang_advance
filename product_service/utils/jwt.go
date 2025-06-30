package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/spf13/viper"
)

// Only needed if you protect POST/PUT/DELETE routes.
func GenerateToken() (string, error) {
    claims := jwt.MapClaims{
        "service": "product",
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
        SignedString([]byte(viper.GetString("JWT_SECRET")))
}
