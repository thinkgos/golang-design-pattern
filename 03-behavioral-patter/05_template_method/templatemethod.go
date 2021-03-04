package templatemethod

import "fmt"

type Downloader interface {
	Download(uri string)
}

// 模板实现接口
type implement interface {
	download()
	save()
}

// 父类模板
type template struct {
	implement
	uri string
}

func newTemplate() *template {
	tpl := &template{}
	tpl.implement = tpl
	return tpl
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) download() {
	fmt.Print("default download\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

/************************** 子类继承实现 ******************/

// 继承父类实现http
type HTTPDownloader struct {
	template
}

func NewHTTPDownloader() Downloader {
	downloader := new(HTTPDownloader)
	downloader.implement = downloader // 让父类调用子类的方法
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

// 继承父类实现ftp
type FTPDownloader struct {
	template
}

func NewFTPDownloader() Downloader {
	downloader := new(FTPDownloader)
	downloader.implement = downloader // 让父类调用子类的方法
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
