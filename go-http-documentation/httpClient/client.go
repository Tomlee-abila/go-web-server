package httpclient

import (
	"fmt"
	"net/http"
)

func GoClient(){
	resp, err := http.Get("https://reqres.in/")

	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(resp)

	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	// resp, err := http.PostForm("http://example.com/form",
	// 	url.Values{"key": {"Value"}, "id": {"123"}})
}