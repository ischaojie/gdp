package dp

import "fmt"

/*模板方法模式*/

// 利用模板方法模式实现一个下载器
// 下载的过程分为：开始--下载--保存--完成，四个步骤
// 其中下载和保存都可以选择具体的实现，比如http下载和ftp下载。
// 通过模板方法定义统一的接口

type Downloader interface {
	Download(url string)
}

// -----------------------------模板------------------------------//

// template 定义了下载器的骨架
// 依次调用
type template struct {
	implement
	url string
}

func (t *template) Download(url string) {
	t.url = url
	fmt.Println("Download start...")
	t.implement.download()
	t.implement.save()
	fmt.Println("Download end...")
}

func (t *template) save() {
	fmt.Println("Default save.")
}

// ----------------------------------------------------------------//

type implement interface {
	download()
	save()
}

// ------------------------HTTP download--------------------------//

type HTTPDownloader struct {
	*template
}

func (d *HTTPDownloader) download() {
	fmt.Printf("Download %s via http.\n", d.url)
}

func (d *HTTPDownloader) save() {
	fmt.Println("http save.")
}

// --------------------------FTP download------------------------//

type FTPDownloader struct {
	*template
}

func (d *FTPDownloader) download() {
	fmt.Printf("Download %s via ftp.\n", d.url)
}

// ftp 使用默认的save方法

// -------------------------------------------------------------- //

func NewTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := NewTemplate(downloader)
	downloader.template = template
	return downloader
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := NewTemplate(downloader)
	downloader.template = template
	return downloader
}
