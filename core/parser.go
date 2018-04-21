package core

import (
	"fmt"
	"net/http"
	// "io"

	// "strings"
	"GoSpider/model"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
	resp *http.Response    // http response
	doc  *goquery.Document // 解析后的dom对象
	// ParaseData *model.ParaseData	// 解析后数据存储的对象
}

// 构造函数
func NewParser() *Parser {
	h := new(Parser)
	return h
}

// 解析html节点
func (h *Parser) HtmlParse(resp *http.Response) error {
	defer resp.Body.Close()
	var err error
	h.doc, err = goquery.NewDocumentFromReader(resp.Body)
	h.resp = resp
	return err
}

// 解析文章内容
func (h *Parser) ArticleParse() (*model.Article, error) {
	article := new(model.Article)
	article.Body = h.doc.Find("div#view2").Text()
	article.Title = h.doc.Find("title").Text()
	article.Url = h.resp.Request.URL.RequestURI()
	return article, nil
}

func (h *Parser) PageParse() (*model.Page, error) {
	page := new(model.Page)
	page.Urls = make([]string, 30)
	h.doc.Find("body .typelist ul").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Find("li a").Attr("href")
		if !exists {
			return
		}
		fullUrl, err := h.resp.Request.URL.Parse(url)
		fmt.Println(fullUrl)
		if err != nil {
			return
		}
		page.Urls = append(page.Urls, fullUrl.String())
	})

	h.doc.Find("body #page a.PageBox").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		if title != "下一页" {
			return
		}

		url, exists := s.Attr("href")
		fmt.Println(url)
		if !exists {
			panic("没有最后一页了")
		}
		fullUrl, err := h.resp.Request.URL.Parse(url)
		fmt.Println(fullUrl)
		if err != nil {
			return
		}
		page.NextPage = fullUrl.String()
	})

	return page, nil
}
