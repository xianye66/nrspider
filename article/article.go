package article

import (

	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)

type InfoArticle struct {
	Title string
	Keywords []string
	Describe string
	From string
	PublicDate string
	Author string
	Content string
}

func (artc *InfoArticle) FindTileByTitle(replaceStr string,doc *goquery.Document){
	title := ""
	titleElement := doc.Find("title")
	if titleElement != nil{
		title = titleElement.Text()
		title = strings.Replace(title,replaceStr,"",0)
	}
	artc.Title = title
}

func (artc *InfoArticle) FindTileByH1(doc *goquery.Document) {
	title := ""
	titleElement := doc.Find("H1").Eq(0)
	if titleElement == nil{
		title = titleElement.Text()
	}
	artc.Title = title
}

func (artc *InfoArticle) FindDescribe(doc *goquery.Document){
	doc.Find("meta").Each(func(index int,item *goquery.Selection) {
		name,exist := item.Attr("name")
		if exist && name == "description"{
			describe,exist :=item.Attr("content")
			if exist{
				artc.Describe = describe
				return
			}

		}
	})
}

func (artc *InfoArticle) FindKeywords(doc *goquery.Document) {
	doc.Find("meta").Each(func(index int,item *goquery.Selection) {
		name,exist := item.Attr("name")
		if exist && name == "keywords"{
			keywordStr,exist :=item.Attr("content")
			if exist{
				artc.Keywords = strings.Split(keywordStr,",")
				return
			}

		}
	})
}

func (artc *InfoArticle) String() string {
	str := fmt.Sprintf(`
	title: %s
	Describe:%s
	keywordlen:%d
	contentlen:%d
	PublicDate:%s
	Author:%s
	From:%s
	 `,artc.Title,artc.Describe,len(artc.Keywords),len(artc.Content),artc.PublicDate,artc.Author,artc.From)
	return str
}

