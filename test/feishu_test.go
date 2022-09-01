package test

import (
	"go-api/config"
	"testing"
)

func TestFs(t *testing.T) {
	config.FsLog().Warn("测试")
}
