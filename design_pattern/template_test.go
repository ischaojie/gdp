package dp

import "testing"

func TestDownloader(t *testing.T) {
	downloader := NewHTTPDownloader()
	downloader.Download("https://baidu.com")
	// Download start...
	// Download https://baidu.com via http.
	// http save.
	// Download end...

	ftpDownloader := NewFTPDownloader()
	ftpDownloader.Download("ftp://baidu.com")
	// Download start...
	// Download ftp://baidu.com via ftp.
	// Default save.
	// Download end...
}
