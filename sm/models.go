package sm

import (
	"encoding/xml"
	"time"
)

// Url представляет собой структуру данных для хранения информации о конкретном URL в карте сайта.
type Url struct {
	Loc        string    `xml:"loc"`                  // Loc содержит адрес URL.
	LastMod    time.Time `xml:"lastmod"`              // LastMod указывает на последнюю дату модификации URL.
	ChangeFreq string    `xml:"changefreq,omitempty"` // ChangeFreq определяет частоту изменения содержимого URL.
	Priority   float32   `xml:"priority,omitempty"`   // Priority устанавливает приоритет URL относительно других URL на вашем сайте.
}

// UrlSet представляет собой структуру данных для хранения множества URL и используется при сериализации и десериализации XML карты сайта.
type UrlSet struct {
	XMLName      xml.Name `xml:"urlset"`                  // XMLName определяет корневой элемент XML.
	Xmlns        string   `xml:"xmlns,attr"`              // Xmlns определяет пространство имен XML для карты сайта.
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`          // XmlnsXsi определяет пространство имен XML Schema instance.
	XsiSchemaLoc string   `xml:"xsi:schemaLocation,attr"` // XsiSchemaLoc указывает на расположение схемы XML.

	Urls []Url `xml:"url"` // Urls содержит срез URL, которые следует включить в карту сайта.
}
