package sm

import (
	"encoding/xml"
	"time"
)

type Url struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq,omitempty"`
	Priority   float32   `xml:"priority,omitempty"`
}

type UrlSet struct {
	XMLName      xml.Name `xml:"urlset"`
	Xmlns        string   `xml:"xmlns,attr"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XsiSchemaLoc string   `xml:"xsi:schemaLocation,attr"`

	Urls []Url `xml:"url"`
}
