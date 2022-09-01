package test

import (
	"github.com/go-playground/assert/v2"
	"go-api/internal/util"
	"testing"
)

func TestFileWalk(t *testing.T) {
	var compare = map[string]string{"login": "user"}
	files, _ := util.FileWalk("../internal/model")
	assert.Equal(t, files, compare)
}
