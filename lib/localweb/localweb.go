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
func Get(url string) (LocalWeb, error) {
	reader, err := os.Open(util.GetFilePath(url))
	if err != nil {
		fmt.Printf("Error when reading file: %s\n", err.Error())
		return LocalWeb{}, err
	}
	body, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Printf("Error when preocessing html file: %s\n", err.Error())
		return LocalWeb{}, err
	}

	links, _ := body.Find("meta[name='fetcher:total_links']").Attr("content")
	images, _ := body.Find("meta[name='fetcher:total_images']").Attr("content")
	fetchTimeStr, _ := body.Find("meta[name='fetcher:fetch_date']").Attr("content")

	fetchTime, err := time.Parse(time.RFC3339, fetchTimeStr)
	if err != nil {
		fmt.Printf("Error when retrieving fetch date: %s\n", err.Error())
		return LocalWeb{}, err
	}
	totalLinks, err := strconv.Atoi(links)
	if err != nil {
		fmt.Printf("Error when retrieving total links: %s\n", err.Error())
		return LocalWeb{}, err
	}
	totalImages, err := strconv.Atoi(images)
	if err != nil {
		fmt.Printf("Error when retrieving total images: %s\n", err.Error())
		return LocalWeb{}, err
	}

	return LocalWeb{
		site:        util.GetUrlName(url),
		body:        body,
		totalLinks:  totalLinks,
		totalImages: totalImages,
		lastFetch:   fetchTime,
	}, nil
}

func (data LocalWeb) PrintMetadata() {
	fmt.Printf("site: %s\n", data.site)
	fmt.Printf("num_links: %d\n", data.totalLinks)
	fmt.Printf("images: %d\n", data.totalImages)
	fmt.Printf("last_fetch: %s\n", data.lastFetch.Format("Mon Jan 02 2006 15:04 MST"))
}
