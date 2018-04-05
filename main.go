package main

import (
	"fmt"
)

func main() {
	u := NewUrlManager()
	h := NewHtmlDownloader()
	u.addNewUrl("https://www.baidu.com/")
	url, err := u.getNewUrl()
	if err != nil {
		panic(err)
	}
	res, err := h.download(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
