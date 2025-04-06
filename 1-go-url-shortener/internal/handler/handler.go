package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-url-shortener/internal/model"
	"go-url-shortener/internal/utils"
)

func CreateShortURL(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ShortenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		shortCode := req.CustomAlias
		if shortCode == "" {
			shortCode = utils.GenerateShortCode(6)
		}

		expiresAt := (*time.Time)(nil)
		if req.ExpiresAt != "" {
			t, err := time.Parse(time.RFC3339, req.ExpiresAt)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration time format"})
				return
			}
			expiresAt = &t
		}

		// Insert into DB
		query := `INSERT INTO urls (original_url, short_code, expires_at) VALUES ($1, $2, $3)`
		_, err := db.Exec(context.Background(), query, req.URL, shortCode, expiresAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
			return
		}

		shortURL := os.Getenv("BASE_URL") + "/" + shortCode
		c.JSON(http.StatusOK, model.ShortenResponse{ShortURL: shortURL})
	}
}

func RedirectURL(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortCode := c.Param("shortcode")

		var originalURL string
		var expiresAt *time.Time

		query := `SELECT original_url, expires_at FROM urls WHERE short_code = $1`
		err := db.QueryRow(context.Background(), query, shortCode).Scan(&originalURL, &expiresAt)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}

		// Check if expired
		if expiresAt != nil && time.Now().After(*expiresAt) {
			c.JSON(http.StatusGone, gin.H{"error": "Short URL has expired"})
			return
		}

		c.Redirect(http.StatusFound, originalURL)
	}
}
