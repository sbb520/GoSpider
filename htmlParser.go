package main

import(
	"github.com/anaskhan96/soup"
	// "fmt"
	"strings"
)

type htmlParser struct {
	// 根url
	// 协议
	// 解析字符串
	domain string
	protocol string
	body string
}

func NewHtmlParser() *htmlParser {
	h := new(htmlParser)
	// h.getUrlDetail(rootUrl)
	return h
}

func (h *htmlParser) getNewUrls (data string) (string, error) {
	// 获取下一个元素的节点
	doc := soup.HTMLParse(data)
	nextNode := doc.Find("li", "class", "next").Find("a")
	nextUrl := nextNode.Attrs()["href"]

	// 获取
	if string(nextUrl[0]) == "/" {
		nextUrl = h.protocol + "://" + h.domain + nextUrl
	}
	return nextUrl, nil
}

func (h *htmlParser) getNewData (data string) (string, error) {
	// 获取下一个元素的节点
	doc := soup.HTMLParse(data)
	node := doc.Find("div", "id", "view2").Find("p")
	// 获取
	body := node.Text()
	return body, nil
}

// 获取domain和protocol
func (h *htmlParser) setUrlDetail(url string) {
	index := strings.Index(url, "//")
	h.protocol = url[:index-1]

	urlOffHttp := url[index+2:]
	end := strings.Index(urlOffHttp, "/")
	h.domain = urlOffHttp[:end]
}
