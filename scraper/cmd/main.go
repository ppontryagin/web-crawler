package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
	//"github.com/anaskhan96/soup"
)

func main() {
	id := 347

	resp, err := soup.Get(fmt.Sprintf("http://www.gp.se/1.%d", id))
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)

	// 1. get all articles in a body container
	articles := doc.
		Find("div", "class", "article__body__richtext container ").
		FindAll("p")
	articlesJoined := ""
	for _, article := range articles {
		articlesJoined += article.Text()
	}
	fmt.Printf("Body:\n%s\n", articlesJoined)

	// 2. get date of the article
	timeDate := doc.Find("time").Attrs()
	fmt.Printf("Dated: %s\n", timeDate["datetime"])

	// 3. get title
	title := doc.Find("title").Text()
	titleSplit := strings.Split(title, "|")[0]
	fmt.Printf("Title: %s\n", titleSplit)

	// 4. get category
	categoryRaw := doc.
		Find("span", "id", fmt.Sprintf("article-data-1.%d", id)) //.
	//Find("category-main")
	//for _, category := range categoryRaw {
	fmt.Printf("Category: %s\n", categoryRaw.Attrs()["category-main"])
	//}
}
