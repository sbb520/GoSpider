# GoSpider
![GoSpider][1]  

go拥有简单的语法以及不错的io以及高并发性能，本人拿go实现一个爬虫项目，进一步的巩固自己go的学习。  
由于是个人兴趣的爬虫项目，捉取一些小网站（你懂）的小说，一方面捉取的内容能让人感到鸡动，一方面这些小网站一般没有什么爬虫访问的限制，能跟专注于爬虫本身的开发过程。

## 开发环境
go：1.6.2  
[goquery][50]: master

其中goquery是一个go的解析html的第三方工具，类似于jquery一样的dom操作，当然你也可以使用go官方的扩展包golang.org/x/net/html来进行解析。



## 运行方法
```
go get github.com/PuerkitoBio/goquery
./test.sh

# 或直接运行二进制文件
./GoSpider
```

## 效果图
![][2]  
![][3]

## 版本
### v1.0
引入爬虫的四大模块：URL管理器，HTML下载器，HTML解析器，持久化模块，由于单纯作为本人测试以及yy，暂不做数据保存...  
此版本的爬虫是简单的循环爬虫模式，不涉及到多线程。

### v2.0
引入多线程爬虫，开启20个下载解析的Goroutines，爬取速度提升接近20倍，小网站的某系列小说3000篇，在大概3分钟左右全部爬取完毕。


[1]: ./doc/img/1.svg
[2]: ./doc/img/run1.jpg
[3]: ./doc/img/run2.jpg

[50]: https://github.com/PuerkitoBio/goquery