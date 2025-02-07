package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(n *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(n.Run())
}
