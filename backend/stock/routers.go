package stock

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func StockRegister(router *gin.RouterGroup, conn *pgx.Conn) {
	router.GET("/", func(c *gin.Context) {
		StockRetrieve(c, conn)
	})
}

func StockRetrieve(c *gin.Context, conn *pgx.Conn) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	total, err := strconv.Atoi(c.DefaultQuery("total", "10"))
	if err != nil || total < 10 {
		total = 10
	}

	stocks, err := GetStocks(c.Request.Context(), conn, page, total)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
}
