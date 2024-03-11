package web

import (
	"fetcher/lib/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
)

type Web struct {
	site        string
	body        *goquery.Document
	totalLinks  int
	totalImages int
	lastFetch   time.Time
}

// Get Web
func New(url string) Web {
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

	return Web{
		site:        filename,
		body:        body,
		totalLinks:  totalLinks,
		totalImages: totalImages,
		lastFetch:   fetchTime,
	}
}

func (data Web) PrintMetadata() {
	fmt.Printf("site: %s\n", data.site)
	fmt.Printf("num_links: %d\n", data.totalLinks)
	fmt.Printf("images: %d\n", data.totalImages)
	fmt.Printf("last_fetch: %s\n", data.lastFetch.Format("Mon Jan 02 2006 15:04 MST"))
}

// Fetch urls and save it to current directory
func Fetch(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go saveWeb(&wg, url)
	}

	wg.Wait()
}

func saveWeb(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	// Get Web
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error when getting web with message: %s", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error url returning response: %d", resp.StatusCode)
		return
	}

	// Process HTML, getting images, links, and append metadata
	docs, err := processHTML(resp)
	if err != nil {
		fmt.Printf("Error when preocessing html file: %s", err.Error())
		return
	}

	// Save Web
	f, err := os.Create(util.GetUrlName(url) + ".html")
	if err != nil {
		fmt.Printf("Error when creating html file: %s", err.Error())
	}
	defer f.Close()

	modifiedHTML, err := goquery.OuterHtml(docs.Selection)
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte(gohtml.Format(modifiedHTML)))
}

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

	// TODO: Save Images and Scripts
	// saveObjects("img", &images)
	// saveObjects("script", &scripts)

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

// func saveObjects(object string, objectMap *map[string]string) {
// 	resp, err := http.Get(url)
// }

// func saveObject()
