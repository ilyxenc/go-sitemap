package sm

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SitemapBuilder struct {
	urlsMap map[string]Url
}

func NewSitemap() *SitemapBuilder {
	return &SitemapBuilder{
		urlsMap: make(map[string]Url),
	}
}

func Read(filePath string) (*SitemapBuilder, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the sitemap file: %w", err)
	}

	var urlset UrlSet
	err = xml.Unmarshal(data, &urlset)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling the urlset: %w", err)
	}

	builder := NewSitemap()
	for _, url := range urlset.Urls {
		builder.Upsert(url)
	}

	return builder, nil
}

func (sb *SitemapBuilder) Get(loc string) (Url, bool) {
	url, ok := sb.urlsMap[loc]
	return url, ok
}

func (sb *SitemapBuilder) Upsert(url Url) {
	sb.urlsMap[url.Loc] = url
}

func (sb *SitemapBuilder) End(filePath string) error {
	urlset := UrlSet{
		Xmlns:        Xmlns,
		XmlnsXsi:     XmlnsXsi,
		XsiSchemaLoc: XsiSchemaLoc,
		Urls:         make([]Url, 0, len(sb.urlsMap)),
	}

	for _, url := range sb.urlsMap {
		urlset.Urls = append(urlset.Urls, url)
	}

	xmlBytes, err := xml.MarshalIndent(urlset, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling the urlset: %w", err)
	}

	xmlHeader := []byte(xml.Header)
	xmlBytes = append(xmlHeader, xmlBytes...)

	err = os.WriteFile(filePath, xmlBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing to the sitemap file: %w", err)
	}

	return nil
}
