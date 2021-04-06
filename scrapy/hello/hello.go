package main

import (
	"net/http"
	"reflect"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main() {
	resp, err := http.Get("https://www.dbooks.org/blazor-webassembly-succinctly-164200202x/")
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
	}
	defer resp.Body.Close()
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if err !=nil {
		fmt.Println(err)
	}
	doc.Find(".col100").Each(func(i int, s *goquery.Selection){
		title := s.Find("H1").Text()
		fmt.Println(i,title)
	})
	body, _ := ioutil.ReadAll(resp.Body)
	body1 := string(body)
	fmt.Printf("%s",reflect.TypeOf(body1))
	fmt.Printf("%s", reflect.TypeOf(body))
}
