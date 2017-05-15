package store

import (
	//"github.com/kevin-zx/go-util/fileUtil"
	"nrspider/article"
	"html/template"
	"github.com/kevin-zx/go-util/errorUtil"
	"os"

	"github.com/snluu/uuid"
)

func unescaped (x string) interface{} { return template.HTML(x) }

func ArticleStoreAsFile(art article.ArticleInfo){
	fileName  :=  uuid.Rand().Hex()+".html"
	f, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	t := template.New("article.html").Funcs(template.FuncMap{"unescaped": unescaped})
	t,err = t.ParseFiles("template/article.html")
	errorUtil.CheckErrorExit(err)

	errorUtil.CheckErrorExit(t.Execute(f,art))
}
