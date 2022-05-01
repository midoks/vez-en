package robot

import (
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"

	"github.com/midoks/vez/internal/lazyregexp"
	"github.com/midoks/vez/internal/mgdb"
)

const (
	DZONE_NAME      = "dzone"
	DZONE_HOME_PAGE = "https://dzone.com"
	LETTER_BYTES    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func isMatchDZONE_Article(url string) bool {
	return lazyregexp.New(`/articles/([\w-]+)`).Regexp().Match([]byte(url))
}

func getMatchDZONE_ID(url string) string {
	m := lazyregexp.New(`/articles/([\w-]+)`).Regexp().FindStringSubmatch(url)
	if len(m) == 2 {
		return m[1]
	}
	return ""
}

func CreateDZONECollector() *colly.Collector {
	dzone := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(2),
		// Attach a debugger to the collector
		// colly.Debugger(&debug.LogDebugger{}),
	)

	dzone.Limit(&colly.LimitRule{
		DomainGlob:  "*dzone.com.*",
		Parallelism: 3,
		Delay:       5 * time.Second,
	})

	// Find and visit all links
	dzone.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if !isMatchDZONE_Article(url) {
			return
		}

		id := getMatchDZONE_ID(url)

		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("%s%s", DZONE_HOME_PAGE, url)
		}

		_, err := mgdb.ContentOriginFindOne(DZONE_NAME, id)
		if err != nil {
			e.Request.Visit(url)
			return
		}
	})

	dzone.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	// Set error handler
	dzone.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "\nError:", err)
	})

	dzone.OnScraped(func(r *colly.Response) {

		url := r.Request.URL.String()
		if isMatchDZONE_Article(url) {

			fmt.Println("match", url)

			doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
			if err != nil {
				return
			}

			id := getMatchDZONE_ID(url)
			contentBody := htmlquery.Find(doc, `//div[@class="content-html"]`)
			contentTitle := htmlquery.Find(doc, `//h1[@class="article-title"]`)

			if len(contentTitle) == 0 {
				return
			}
			if len(contentBody) == 0 {
				return
			}

			title := htmlquery.OutputHTML(contentTitle[0], false)
			html := htmlquery.OutputHTML(contentBody[0], false)

			_, err = mgdb.ContentAdd(mgdb.Content{
				Url:    url,
				Source: DZONE_NAME,
				User:   DZONE_NAME,
				Id:     id,
				Title:  title,
				Html:   html,
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	return dzone
}

func RunDZONE() {

	app := CreateDZONECollector()
	r, err := mgdb.ContentRandSource(DZONE_NAME)
	if err == nil {
		fmt.Println("rand visiting: ", r.Url)
		app.Visit(r.Url)
	} else {
		fmt.Println("visiting start")
		app.Visit(DZONE_HOME_PAGE)
	}

	app.Wait()
}
