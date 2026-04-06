package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

// TestJWTMiddleware_ValidToken verifies if the middleware allows access with a valid JWT token
func TestJWTMiddleware_ValidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"foo": "bar"})
	tokenString, _ := token.SignedString([]byte("testsecret"))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(JWTMiddleware())
	r.GET("/protected", func(c *gin.Context) { c.String(200, "ok") })

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// TestJWTMiddleware_InvalidToken verifies if the middleware blocks access with an invalid JWT token
func TestJWTMiddleware_InvalidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(JWTMiddleware())
	r.GET("/protected", func(c *gin.Context) { c.String(200, "ok") })

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
