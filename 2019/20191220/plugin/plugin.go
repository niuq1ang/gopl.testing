package plugin

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

// CheckPluginURI
func CheckPluginURI(ctx iris.Context) {
	uri, ok := checkPrefix(ctx)
	if ok {
		transferRequest(uri, ctx)

	} else {
		ctx.Next()

	}
}

func checkPrefix(ctx iris.Context) (string, bool) {
	ctx.Next()
	return "", true
}

func checkSuffix(ctx iris.Context) (string, bool) {
	ctx.Next()
	return "", true
}

func transferRequest(uri string, ctx iris.Context) {
	client := &http.Client{}
	bodyReaderCloser, err := ctx.Request().GetBody()
	if err != nil {

	}
	req, err := http.NewRequest(ctx.Method(), uri, bodyReaderCloser)
	//注： 设置Request头部信息
	for k, v := range ctx.Request().Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//注： 设置Response头部信息
	for k, v := range resp.Header {
		for _, vv := range v {
			ctx.ResponseWriter().Header().Set(k, vv)
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(data)
	}

}
