package test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"go-api/pkg/jwt"
	"testing"
)

func TestAesCtrCrypt(t *testing.T) {
	source := "hello world"
	key := "key"
	fmt.Println("原字符：", source)
	encryptCode, _ := jwt.AesCtrCrypt([]byte(source), key)
	fmt.Println("密文：", string(encryptCode))

	decryptCode, _ := jwt.AesCtrCrypt(encryptCode, key)

	fmt.Println("解密：", string(decryptCode))
	assert.Equal(t, source, string(decryptCode))
}
