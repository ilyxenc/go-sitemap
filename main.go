package main

import (
	"fmt"
	"time"

	"github.com/ilyxenc/go-sitemap/sm"
)

func main() {
	sitemap := sm.NewSitemap()

	sitemap.Upsert(sm.Url{
		Loc:        "https://example.com/",
		LastMod:    time.Now(),
		ChangeFreq: "daily",
		Priority:   1.0,
	})

	if err := sitemap.End("sitemap.xml"); err != nil {
		fmt.Println("Error:", err)
	}

	sitemap2, err := sm.Read("sitemap.xml")
	if err != nil {
		fmt.Println("Error:", err)
	}

	sitemap2.Upsert(sm.Url{
		Loc:     "https://example.com/",
		LastMod: time.Now(),
	})

	if err := sitemap2.End("sitemap.xml"); err != nil {
		fmt.Println("Error:", err)
	}

	url, ok := sitemap.Get("https://example.com/")
	if !ok {
		fmt.Println("URL not found")
	} else {
		fmt.Println(url)
	}
}
