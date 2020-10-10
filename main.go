package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	getPages()
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination")
	return 0

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request falled with Status: ", res.StatusCode)
	}
}
