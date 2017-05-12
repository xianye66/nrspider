package extender

import (
	"github.com/PuerkitoBio/gocrawl"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"nrspider/spidertools"
)

var urlMap = make(map[string]int)

type MyExtender struct {
	gocrawl.DefaultExtender
	SpiderTool spidertools.SpiderTool

}

func (mextd *MyExtender) New(spiderTool spidertools.SpiderTool) {
	mextd.SpiderTool = spiderTool
}

func (mextd *MyExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (harvested interface{}, findLinks bool)  {
	mextd.SpiderTool.Extract(doc,*ctx.URL())
	return nil,true
}


// Override Filter for our need.
func (mextd *MyExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	println(ctx.URL().String())
	if urlMap[ctx.URL().String()] == 1 {
		return false
	}else{
		urlMap[ctx.URL().String()] = 1
		return mextd.SpiderTool.UrlRule(*ctx.URL())
	}
}