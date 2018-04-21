package main

import (
	"GoSpider/core"
	"fmt"
)

func main() {
	rootUrl := "http://www.18yyyy.com/t01/index.html"

	urlManager := core.NewUrlManager()

	// 导入根url
	urlManager.AddNewUrl(rootUrl)

	//
	artList := make(chan string, 20)

	// 新建20个goroutines来读取每页的文章
	for i := 0; i < 20; i++ {
		go func() {
			for list := range artList {
				// 下载url的内容
				downloader := core.NewDownloader()
				response, err := downloader.Download(list)
				if err != nil {
					panic(err)
				}

				// 解析页面内容
				parser := core.NewParser()
				parser.HtmlParse(response)
				page, err := parser.ArticleParse()
				fmt.Println(page.Body)
			}

		}()
	}

	// 导出每个目录页
	for {
		url, err := urlManager.GetNewUrl()
		if err != nil {
			panic(err)
		}

		// 下载url的内容
		downloader := core.NewDownloader()
		response, err := downloader.Download(url)

		// 解析页面内容
		fmt.Println(url)
		parser := core.NewParser()
		parser.HtmlParse(response)
		page, err := parser.PageParse()
		if err != nil {
			panic(err)
		}
		for _, url := range page.Urls {
			if url != "" {
				artList <- url
			}
		}
		urlManager.AddNewUrl(page.NextPage)
		//
	}

}
