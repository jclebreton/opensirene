package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUI returns the UI
func (v *ViewsContext) GetUI(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// GetUIData returns UI data
func (v *ViewsContext) GetUIData(c *gin.Context) {
	data := c.Param("data")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"data": data})
}
