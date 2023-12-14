package main

import (
    "flag"
    "fmt"
    "github.com/gocolly/colly"
    "strings"
)

func main() {

    maxElements := flag.Int("max", 10, "Max number of comments to show")

	flag.Parse()

    c := colly.NewCollector(
        colly.AllowedDomains("www.ptt.cc"),
    )

	count := 1

    c.OnHTML(".push", func(e *colly.HTMLElement) {
		if count > *maxElements {
            return
        }

        userName := e.ChildText(".push-userid")
        comment := strings.TrimSpace(e.ChildText(".push-content"))
        timeStamp := strings.TrimSpace(e.ChildText(".push-ipdatetime"))

        fmt.Printf("%d. 名字：%s，留言%s，時間： %s\n", count, userName, comment, timeStamp)

		count++
    })

    // c.OnRequest(func(r *colly.Request) {
    //     fmt.Println("Visiting", r.URL)
    // })

    c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
}
