package file

import (
	"os"
	"path/filepath"
	"strings"
)

// Put 将数据读入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileNameWithoutExtension 去除文件后缀
func FileNameWithoutExtension(name string) string {
	return strings.TrimSuffix(name, filepath.Ext(name))
}
