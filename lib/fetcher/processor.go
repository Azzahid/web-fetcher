package fetcher

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Process HTML by
// Getting images, links, and append metadata
// Replacing images with local mirror
// Replacing scripts with local mirror
// Replacing css scripts with local mirror
func processHTML(resp *http.Response) (*goquery.Document, error) {
	body, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	//Get Images, URLs and Scripts
	urls := getHTMLObjects(body, "a", "href")
	images := getHTMLObjects(body, "img", "src")
	// scripts := getHTMLObjects(body, "script", "src")
	// css := getHTMLObjects(body, "link[rel='stylesheet']", "src")

	// // TODO: Save Images and Scripts
	// saveObjects(&images)
	// saveObjects(&scripts)
	// saveObjects(&css)

	//TODO: Redirect Urls and Scripts in the body

	//Append metadata
	appendMetadata(body, len(urls), len(images))

	return body, nil
}

func appendMetadata(docs *goquery.Document, totalUrls, totalImages int) {
	head := docs.Find("head")
	metaBase := "<meta name=\"%s\" content=\"%s\">"

	head.AppendHtml(fmt.Sprintf(metaBase, "fetcher:total_links", strconv.Itoa(totalUrls)))
	head.AppendHtml(fmt.Sprintf(metaBase, "fetcher:total_images", strconv.Itoa(totalImages)))
	head.AppendHtml(fmt.Sprintf(metaBase, "fetcher:fetch_date", time.Now().UTC().Format(time.RFC3339)))
}

func getHTMLObjects(body *goquery.Document, object string, refTag string) map[string]string {
	results := make(map[string]string)
	body.Find(object).Each(func(_ int, s *goquery.Selection) {
		str, ok := s.Attr(refTag)
		if ok {
			results[str] = str
		}
	})

	return results
}

// func saveObjects(objectMap *map[string]string) {
// 	var wg sync.WaitGroup

// 	for i, url := range objectMap {
// 		wg.Add(1)
// 		go saveObject(wg, url)
// 	}

// 	wg.Wait()
// }

// func saveObject(wg *sync.WaitGroup, url string) {

// }
