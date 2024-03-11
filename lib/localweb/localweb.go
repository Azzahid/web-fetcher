package localweb

import (
	"fetcher/lib/util"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type LocalWeb struct {
	site        string
	body        *goquery.Document
	totalLinks  int
	totalImages int
	lastFetch   time.Time
}

// Get Web
func New(url string) LocalWeb {
	filename := util.GetUrlName(url)
	reader, err := os.Open(filename + ".html")
	if err != nil {
		fmt.Printf("Error when reading file: %s\n", err.Error())
		panic(err.Error())
	}
	body, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Printf("Error when preocessing html file: %s", err.Error())
		panic(err.Error())
	}

	links, _ := body.Find("meta[name='fetcher:total_links']").Attr("content")
	images, _ := body.Find("meta[name='fetcher:total_images']").Attr("content")
	fetchTimeStr, _ := body.Find("meta[name='fetcher:fetch_date']").Attr("content")

	fetchTime, err := time.Parse(time.RFC3339, fetchTimeStr)
	if err != nil {
		fmt.Printf("Error when retrieving fetch date: %s", err.Error())
		panic(err.Error())
	}
	totalLinks, err := strconv.Atoi(links)
	if err != nil {
		fmt.Printf("Error when retrieving total links: %s", err.Error())
		panic(err.Error())
	}
	totalImages, err := strconv.Atoi(images)
	if err != nil {
		fmt.Printf("Error when retrieving total images: %s", err.Error())
		panic(err.Error())
	}

	return LocalWeb{
		site:        filename,
		body:        body,
		totalLinks:  totalLinks,
		totalImages: totalImages,
		lastFetch:   fetchTime,
	}
}

func (data LocalWeb) PrintMetadata() {
	fmt.Printf("site: %s\n", data.site)
	fmt.Printf("num_links: %d\n", data.totalLinks)
	fmt.Printf("images: %d\n", data.totalImages)
	fmt.Printf("last_fetch: %s\n", data.lastFetch.Format("Mon Jan 02 2006 15:04 MST"))
}
