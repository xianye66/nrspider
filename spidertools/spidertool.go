package spidertools

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)
type SpiderTool interface {
	Extract(doc *goquery.Document, url url.URL)
	UrlRule(url url.URL) bool

}
type DefaultSpiderTool struct {
	SpiderTool
}

func (s *DefaultSpiderTool) Extract(doc *goquery.Document, url url.URL){
	if strings.Contains(url.String(), ".html") {
		extractContent(doc)
	}
}


func extractContent(doc *goquery.Document)  {

}

func (s *DefaultSpiderTool) UrlRule(url url.URL) bool {
	return true
}

func (s *DefaultSpiderTool) DecodeToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}

