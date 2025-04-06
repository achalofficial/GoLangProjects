package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(r *gin.Engine, db *pgxpool.Pool) {
	api := r.Group("/api")
	{
		api.POST("/shorten", CreateShortURL(db))
	}

	r.GET("/:shortcode", RedirectURL(db))
}
