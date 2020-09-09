package main

import (
	"fmt"
	"net/http"

	"go.elastic.co/apm/module/apmhttp"
)

func main() {
	fmt.Println("start ...")
	mux := http.NewServeMux()
	// ...
	mux.HandleFunc("/apm/client/login", Login)
	http.ListenAndServe(":8200", apmhttp.Wrap(mux))
	fmt.Println("我裂开了 ...")

}

func Login(reponse http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(reponse, string("你怎么就像登录啊"))
}
