package main

import (
	"github.com/PuerkitoBio/gocrawl"
	"nrspider/extender"
	"time"
	"nrspider/spidertools"
)

func main(){
	extender := extender.MyExtender{SpiderTool: &spidertools.YiCaiTool{}}
	opts := gocrawl.NewOptions(&extender)

	opts.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36"
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	//opts.MaxVisits = 2
	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("http://www.yicai.com/")
	
}