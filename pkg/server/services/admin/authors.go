package admin

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	// "net/url"
	"database/sql"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/golang-jwt/jwt"

	// "github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

func (s *AdminPageService) LoginAuthReqHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusBadRequest, "Missing Authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("your_secret_key"), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, usernameOk := claims["username"].(string)
		password, passwordOk := claims["password"].(string)
		if !usernameOk || !passwordOk {
			return c.JSON(http.StatusBadRequest, "Invalid token payload")
		}

		// fmt.Printf("Extracted Email: %s, Password: %s", username, password)

		var req db.CreateAuthorParams
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
		}

		params := db.GetAuthorByUsernameAndPasswordParams{
			Name:         username,
			PasswordHash: password,
		}

		resp, err := s.Queries.GetAuthorByUsernameAndPassword(c.Request().Context(), params)

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, "Invalid email or password")
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, "some error")
		}

		newClaims := jwt.MapClaims{
			"user": map[string]interface{}{
				"id":       resp.ID,
				"email":    resp.Name,
				"password": resp.PasswordHash,
				"roles":    req.Role,
			},
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

		signedToken, err := newToken.SignedString([]byte("your_secret_key"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to generate token")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": signedToken,
		})
	}

	return c.JSON(http.StatusUnauthorized, "Invalid token")
}

func (s *AdminPageService) CreateNewAuthor(c echo.Context) error {
	var req db.CreateAuthorParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	blog, err := s.Queries.CreateAuthor(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create author"})
	}

	return c.JSON(http.StatusOK, blog)
}

func (s *AdminPageService) GetAllAuthors(c echo.Context) error {
	authors, err := s.Queries.GetAuthors(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve authors"})
	}

	return c.JSON(http.StatusOK, authors)
}
