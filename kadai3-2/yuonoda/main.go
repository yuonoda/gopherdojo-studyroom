package main

import (
	"flag"
	downloader "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/downloader"
)

var url = flag.String("url", "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2", "URL to download")
var batchCount = flag.Int("c", 1, "how many times you request content")
var dwDirPath = flag.String("path", ".", "where to put a downloaded file")

func main() {
	flag.Parse()
	d := downloader.Downloader{Url: *url}
	d.Download(*batchCount, *dwDirPath)
}
