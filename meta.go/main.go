package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	ans, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(ans)

}
