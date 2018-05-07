package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func parsePagination(c *gin.Context, pg *Pagination) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return err
	}

	pg.Page = page
	pg.Size = size
	return nil
}