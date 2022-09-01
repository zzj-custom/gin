package util

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetRequestParams(ctx *gin.Context, method string) string {
	var params string
	if method == http.MethodDelete || method == http.MethodPost || method == http.MethodPatch || method == http.MethodHead {
		postForm := ctx.Request.PostForm
		if len(postForm) == 0 {
			// 解决io读取之后，后续操作读取不到参数的问题
			body, err := ctx.GetRawData()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("data: %v\n", string(body))
			//把读过的字节流重新放到body
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			params = string(body)
		} else {
			params = postForm.Encode()
		}
	} else if method == http.MethodGet {
		params = ctx.Request.URL.RawQuery
	}
	return params
}
