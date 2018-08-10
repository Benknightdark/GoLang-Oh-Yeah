package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	// c.OnHTML("script", func(e *colly.HTMLElement) {

	// 	e.Request.Visit(e.Text)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// fmt.Println(r.ResponseCharacterEncoding)
	// if strings.Contains(string(r.Body.), "window._sharedData") == true {
	// 	fmt.Println("Visiting", r)
	// }
	//fmt.Println(r)

	// })
	c.OnResponse(func(r *colly.Response) {
		re := regexp.MustCompile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png)`)
		tt := string(r.Body)
		match := re.FindStringSubmatch(tt)
		fmt.Println(match)
		//fmt.Println(string(r.Body))
	})

	c.Visit("https://www.instagram.com/kat.mcnamara/")
}
