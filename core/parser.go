package core

import(
	// "github.com/anaskhan96/soup"
	"io"
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"GoSpider/model"
)

type parser struct {
	// 根url
	// 协议
	// 解析字符串
	// 解析的内容
	domain string
	protocol string
	body string
	url string
	doc *goquery.Document
	ParaseData *model.ParaseData
}

// 构造函数
func NewParser() *parser {
	h := new(parser)
	h.ParaseData = new(model.ParaseData)
	
	return h
}

func (h *parser) Init(data io.ReadCloser, url string) {
	h.url = url
	h.HtmlParse(data) // 解析url下的数据
	h.SetUrlDetail(url) // 解析url的域名

	h.SetNewData() //
	h.SetNewUrls()
	h.SetTitle()
	h.SetUrl()
}

// 解析html节点
func (h *parser) HtmlParse(data io.ReadCloser) error{
	defer data.Close()
	var err error
	h.doc, err = goquery.NewDocumentFromReader(data)
	
	return err
}

// 获取domain和protocol
func (h *parser) SetUrlDetail(url string) error{
	index := strings.Index(url, "//")
	h.protocol = url[:index-1]

	urlOffHttp := url[index+2:]
	end := strings.Index(urlOffHttp, "/")
	h.domain = urlOffHttp[:end]

	return nil
}

// 解析出下一个新的url
func (h *parser) SetNewUrls () (error) {

	nextUrl, err := h.doc.Find("li.next").Find("a").Attr("href")
	fmt.Println(err)

	if string(nextUrl[0]) == "/" {
		nextUrl = h.protocol + "://" + h.domain + nextUrl
	} 
	h.ParaseData.NextUrl = nextUrl
	return nil
}

// 解析出数据
func (h *parser) SetNewData () (error) {
	body := h.doc.Find("div#view2").Text()
	h.ParaseData.Data = body
	return nil
}

// 解析出当前的标题
func (h *parser) SetTitle() error{
	title := h.doc.Find("title").Text()
	h.ParaseData.Title = title
	
	return nil
}

// 
func (h *parser) SetUrl() {
	h.ParaseData.Url = h.url
}

// 获取解析后的数据
func (h *parser) GetParaseData() *model.ParaseData{
	return h.ParaseData
}






