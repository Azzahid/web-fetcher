package fetcher

import (
	"fetcher/lib/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
)

// Fetch urls and save it to current directory
func Fetch(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go saveWeb(&wg, url)
	}

	wg.Wait()
}

// Save web into current directory
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
