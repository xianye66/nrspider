package spidertools

import (
	"github.com/PuerkitoBio/goquery"
	"nrspider/article"
	"strings"
	"net/url"
	"nrspider/store"

)

type YiCaiTool struct {
	DefaultSpiderTool
}

func (yct *YiCaiTool) Extract(doc *goquery.Document, url url.URL)  {
	if !strings.Contains(url.String(),".html") {
		return
	}
	atc := new(article.ArticleInfo)
	atc.FindTileByTitle("_第一财经",doc)
	atc.FindDescribe(doc)
	atc.FindKeywords(doc)

	//atc.FindTileByH1(doc)
	atc.Author = yct.findAuthor(doc)
	atc.PublicDate = yct.findPublicDate(doc)
	atc.Content = yct.findContent(doc)
	atc.From = "第一财经"
	println(atc.String())
	store.Store_article(*atc)
	println(1)
}

func (yct *YiCaiTool) findContent(doc *goquery.Document) string {

	contentHtml := ""
	contentContanier := doc.Find("div.m-text")
	if contentContanier != nil {
		contentHtml,_ = contentContanier.Html()
	}
	return contentHtml
}

func (s *YiCaiTool) UrlRule(url url.URL) bool {
	println(url.String())
	if strings.Contains(url.String(),".html") || url.String() =="http://www.yicai.com/"{
		return true
	}else{
		return false
	}
}

func (yct *YiCaiTool) findAuthor(doc *goquery.Document) string  {
	author := ""
	authorElement := doc.Find("h3.f-ff1.f-fwn.f-fs14")
	authorElement.Each(func(index int,item *goquery.Selection) {
		if strings.Contains(item.Text(),"编辑"){
			author = strings.Replace(item.Text(),"编辑：","",1)
		}
	})

	return author
}

func (yct *YiCaiTool) findPublicDate(doc *goquery.Document) string  {
	publicDate := ""
	publicDateElemnt := doc.Find("H2.f-ff1.f-fwn.f-fs14 span")
	if publicDateElemnt != nil {
		publicDateElemnt.Each(func(index int,sel *goquery.Selection) {
			if sel.Text() != "" {
				publicDate = sel.Text()
			}
		})
	}
	return publicDate
}