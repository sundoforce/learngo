package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	totalPages := getPages()

	for i := 9; i < totalPages; i++ {
		getPage(i)
	}
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func getPage(page int) {
	pageURL := baseURL + "%start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
}
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("a").Length())
	})
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
