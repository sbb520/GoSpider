package main

import (
	"fmt"
)

func main() {
	rootUrl := "http://www.18yyyy.com/htm/2018/4/4/t01/404695.html"


	urlManager := NewUrlManager()
	htmlDownloader := NewHtmlDownloader()
	htmlParser := NewHtmlParser()

	// 导入根url
	urlManager.addNewUrl(rootUrl)

	for i := 0; i < 2; i++ {
		url, err :=urlManager.getNewUrl()
		if err != nil {
			panic(err)
		}
		// 下载url里的内容
		res, err := htmlDownloader.download(url)
		if err != nil {
			panic(err)
		}
		fmt.Println(url)
		htmlParser.setUrlDetail(url)
		// 解析网页内容
		newUrl, _ := htmlParser.getNewUrls(res)
		body, _ := htmlParser.getNewData(res)
		urlManager.addNewUrl(newUrl)
		fmt.Println(body)
		break
	}
	
	
}
