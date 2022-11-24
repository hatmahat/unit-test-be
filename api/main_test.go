package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // ubah ke mode test biar ngga rame pas tes
	os.Exit(m.Run())
}
