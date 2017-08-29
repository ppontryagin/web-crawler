package scraper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

// ScrapPage returns date,title,category,body for a giver gp id
func ScrapPage(id int) (link, date, title, category, body string, err error) {
	link = fmt.Sprintf("http://www.gp.se/1.%d", id)
	fmt.Printf("Dumping link %s\n", link)
	resp, err := soup.Get(link)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)

	// 1. get all articles in a body container
	articlesWithP, err := doc.
		Find("div", "class", "article__body__richtext container ")
	if err != nil {
		return "", "", "", "", "", errors.New("empty web page")
	}
	articles := articlesWithP.FindAll("p")
	if len(articles) == 0 {
		return "", "", "", "", "", errors.New("not a valid web page")
	}
	for _, article := range articles {
		// TODO:analyse other tag for text
		// iterate over <strong> tags in <p>
		artNoTag, err := article.Find("strong") // bold
		if err == nil {
			body += artNoTag.Text()
		} else {
			artNoTag, err = article.Find("em") // italic
			if err == nil {
				body += artNoTag.Text()
			} else { // plain text
				body += article.Text()
			}
		}
	}

	// remove | from the body
	body = strings.Replace(body, "|", "", -1)

	// 2. get date of the article
	timeDateRawAttrs, err := doc.Find("time")
	if err != nil {
		return "", "", "", "", "", errors.New("date cannot be found")
	}
	timeDateRaw := timeDateRawAttrs.Attrs()

	date = timeDateRaw["datetime"]

	// 3. get title
	titleRawText, err := doc.Find("title")
	if err != nil {
		return "", "", "", "", "", errors.New("title cannot be found")
	}
	titleRaw := titleRawText.Text()
	title = strings.Split(titleRaw, "|")[0]
	//fmt.Println("title, ", title)

	// 4. get category
	categoryRaw, err := doc.Find("span", "id", fmt.Sprintf("article-data-1.%d", id))
	if err != nil {
		return "", "", "", "", "", errors.New("category cannot be found")
	}
	category = categoryRaw.Attrs()["category-main"]

	return
}
