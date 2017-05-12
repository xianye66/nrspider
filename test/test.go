package main

import (
	"regexp"
	"github.com/PuerkitoBio/gocrawl"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
)

// Only enqueue the root and paths beginning with an "a"
var rxOk = regexp.MustCompile(`http://www.baidu\.com.*$`)

// Create the MyExtender implementation, based on the gocrawl-provided DefaultExtender,
// because we don't want/need to override all methods.
type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visit and Filter
}

// Override Visit for our need.
func (x *ExampleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// Use the goquery document or res.Body to manipulate the data
	// ...
	println("111111111111111111111111111111111111111111111")
	println(ctx.URL())
	s,_ := doc.Html()
	println(s)
	// Return nil and true - let gocrawl find the links
	return nil, true
}

// Override Filter for our need.
func (x *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	println("-------------------------------------"+ctx.NormalizedURL().String())
	println("-------------------------------------"+ctx.URL().String())
	return true
	//return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

func ExampleCrawl() {
	// Set custom options
	opts := gocrawl.NewOptions(new(ExampleExtender))

	// should always set your robot name so that it looks for the most
	// specific rules possible in robots.txt.
	//opts.RobotUserAgent = "Example"
	// and reflect that in the user-agent string used to make requests,
	// ideally with a link so site owners can contact you if there's an issue
	opts.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36"

	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll

	// Play nice with ddgo when running the test!
	opts.MaxVisits = 2

	// Create crawler and start at root of duckduckgo
	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("https://www.baidu.com/")

	// Remove "x" before Output: to activate the example (will run on go test)

	// xOutput: voluntarily fail to see log output
}

func main()  {
	ExampleCrawl()
}