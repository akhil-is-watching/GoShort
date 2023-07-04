package api

import (
	"net/http"

	"github.com/akhil-is-watching/goshort/types"
	"github.com/gin-gonic/gin"
)

func (s *APIServer) CreateShort(c *gin.Context) {
	var input types.GoShortInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	short, err := s.store.CreateShort(&input)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"short": short})
}
