package util

import (
	"os"
	"path/filepath"
	"strings"
)

// FileWalk 获取某个目录下面的所有子文件
func FileWalk(path string) (map[string]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), "model.") {
			dir := filepath.Dir(path)
			index := strings.LastIndex(dir, "/")
			dirname := dir[index+1:]
			files[dirname] = strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		}
		return nil
	})
	return files, err
}
