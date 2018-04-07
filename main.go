package main

import (
	"fmt"
	"GoSpider/core"
)

func main() {
	rootUrl := "http://www.18yyyy.com/htm/2018/4/4/t01/404693.html"

	urlManager := core.NewUrlManager()
	downloader := core.NewDownloader()
	parser := core.NewParser()

	// 导入根url
	urlManager.AddNewUrl(rootUrl)

	for i := 0; i < 2; i++ {
		url, err :=urlManager.GetNewUrl()
		if err != nil {
			panic(err)
		}
		// 下载url里的内容
		body, err := downloader.Download(url)
		if err != nil {
			panic(err)
		}
		// 解析数据
		parser.Init(body, url)
		data := parser.GetParaseData()
		// 添加下一页的url
		urlManager.AddNewUrl(data.NextUrl)
		// 
		fmt.Println(data.Data)
	}
	
	
}
