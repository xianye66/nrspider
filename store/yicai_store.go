package store

import (
	//"github.com/kevin-zx/go-util/fileUtil"
	"nrspider/article"
	"html/template"
	"github.com/kevin-zx/go-util/errorUtil"
	"os"
)

func unescaped (x string) interface{} { return template.HTML(x) }

func Store_article(art article.ArticleInfo){
	t := template.New("article.html").Funcs(template.FuncMap{"test": unescaped})
	t,err := t.ParseFiles("template/article.html")
	errorUtil.CheckErrorExit(err)
	errorUtil.CheckErrorExit(t.Execute(os.Stdout,art))
}
