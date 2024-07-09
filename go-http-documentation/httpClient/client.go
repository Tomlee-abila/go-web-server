package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GoClient(){
	resp, err := http.Get("https://reqres.in")

	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	// resp, err := http.PostForm("http://example.com/form",
	// 	url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil{
		fmt.Println("Error!\n\n",err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil{
		fmt.Println("Error!\n\n",err)
		return
	}

	os.WriteFile("index.html", body, 0064)

	fmt.Println("Status OK!!\n\n",string(body))

	
}