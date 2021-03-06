package api

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

func TestProductsTextSearch(t *testing.T) {
	if err := repositories.InitDB(true); err != nil {
		t.Error("An error has ocurred", err)
	}

	defer ResetDB()

	r := gin.New()

	method := "GET"
	path := "/api/search/products"

	repositories.Spree_db.Create(&domain.User{})

	r.GET(path, func(c *gin.Context) {
		user := &domain.User{}
		repositories.Spree_db.First(user)
		c.Set("CurrentUser", user)
		ProductsIndex(c)
	})
	w := PerformRequest(r, method, path, nil)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d, but was %d", http.StatusOK, w.Code)
	}
}
