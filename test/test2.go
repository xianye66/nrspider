package main

import (
	"log"
	"net/url"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
)

func main()  {
	saveImages("http://studygolang.com/static/img/logo1.png")
}

func saveImages(img_url string){
	log.Println(img_url)
	u, err := url.Parse(img_url)
	if err != nil {
		log.Println("parse url failed:", img_url, err)
		return
	}

	//去掉最左边的'/'
	tmp := strings.TrimLeft(u.Path, "/")
	filename := "./"+strings.ToLower(strings.Replace(tmp, "/", "-", -1))

	exists := checkExists(filename)
	if exists {
		return
	}

	response, err := http.Get(img_url)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", img_url, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	defer image.Close()
	image.Write(data)
}


func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}