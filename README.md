# GoSpider
![GoSpider][1]  

go拥有简单的语法以及不错的io以及高并发性能，本人拿go实现一个爬虫项目，进一步的巩固自己go的学习。  
由于是个人兴趣的爬虫项目，捉取一些小网站（你懂）的小说，一方面捉取的内容能让人感到鸡动，一方面这些小网站一般没有什么爬虫访问的限制，能跟专注于爬虫本身的开发过程。

## 开发环境
go：1.6.2  
[goquery][50]: master

其中goquery是一个go的解析html的第三方工具，类似于jquery一样的dom操作，当然你也可以使用go官方的扩展包golang.org/x/net/html来进行解析。

[1]: ./doc/img/1.svg

[50]: https://github.com/PuerkitoBio/goquery

## 运行方式
```
./test.sh
```

## 版本
### v1.0
引入爬虫的四大模块：URL管理器，HTML下载器，HTML解析器，持久化模块。
此版本的爬虫是简单的循环爬虫模式，不涉及到多线程。