package main

import (
	"html/template"
	"os"
)
func main()  {
	t := template.New("").Funcs(template.FuncMap{"test": unescaped})
	println(t==nil)
	t,_ = template.ParseFiles("template/article.html")
	println(t==nil)
	//t = t.Funcs(template.FuncMap{"test": unescaped})
	//if err != nil{
	//	println(err.Error())
	//}
	data := struct {
		Title string
		KeywordsStr string
		Content string
		From string
		Author string
		PublicDate string
	}{
		Title:"title",
		KeywordsStr:"ks",
		Content:"content",
		From:"from",
		Author:"Author",
		PublicDate:"pd",
	}
	println(t==nil)
	err := t.Execute(os.Stdout, data)
	if err != nil{
		println(err.Error())
	}

}


func unescaped (x string) interface{} { return template.HTML(x) }