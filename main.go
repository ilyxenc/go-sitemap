package main

import (
	"fmt"
	"time"

	"github.com/ilyxenc/go-sitemap/sm"
)

func main() {
	sitemap := sm.NewSitemap()

	sitemap.Upsert(sm.Url{
		Loc:        "http://example.com",
		LastMod:    time.Now(),
		ChangeFreq: "daily",
		Priority:   1,
	})

	if err := sitemap.End("path/to/output.xml"); err != nil {
		fmt.Println("Error: ", err)
	}
}
