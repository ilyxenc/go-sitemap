package sm

import (
	"encoding/xml"
	"fmt"
	"os"
)

// SitemapBuilder структура для создания и обработки карты сайта.
type SitemapBuilder struct {
	urlsMap map[string]Url
}

// NewSitemap создает новый экземпляр SitemapBuilder.
func NewSitemap() *SitemapBuilder {
	return &SitemapBuilder{
		urlsMap: make(map[string]Url),
	}
}

// Read считывает файл карты сайта из указанного пути и возвращает экземпляр SitemapBuilder.
func Read(filePath string) (*SitemapBuilder, error) {
	// Чтение содержимого файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the sitemap file: %w", err)
	}

	// Разбор XML-данных в структуру UrlSet
	var urlset UrlSet
	err = xml.Unmarshal(data, &urlset)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling the urlset: %w", err)
	}

	// Создание нового экземпляра SitemapBuilder и заполнение его данными из UrlSet
	builder := NewSitemap()
	for _, url := range urlset.Urls {
		builder.Upsert(url)
	}

	return builder, nil
}

// Get возвращает URL по указанному местоположению (loc).
func (sb *SitemapBuilder) Get(loc string) (Url, bool) {
	url, ok := sb.urlsMap[loc]
	return url, ok
}

// Upsert обновляет или добавляет URL в urlsMap.
func (sb *SitemapBuilder) Upsert(url Url) {
	sb.urlsMap[url.Loc] = url
}

// End завершает создание карты сайта и записывает ее в указанный файл.
func (sb *SitemapBuilder) End(filePath string) error {
	// Создание нового экземпляра UrlSet с необходимыми значениями из констант и объекта SitemapBuilder
	urlset := UrlSet{
		Xmlns:        Xmlns,
		XmlnsXsi:     XmlnsXsi,
		XsiSchemaLoc: XsiSchemaLoc,
		Urls:         make([]Url, 0, len(sb.urlsMap)),
	}

	// Заполнение UrlSet данными из SitemapBuilder
	for _, url := range sb.urlsMap {
		urlset.Urls = append(urlset.Urls, url)
	}

	// Размещение UrlSet в формат XML с отступами
	xmlBytes, err := xml.MarshalIndent(urlset, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling the urlset: %w", err)
	}

	// Добавление заголовка XML и запись данных в файл
	xmlHeader := []byte(xml.Header)
	xmlBytes = append(xmlHeader, xmlBytes...)

	err = os.WriteFile(filePath, xmlBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing to the sitemap file: %w", err)
	}

	// Возвращение nil, если операции завершились успешно
	return nil
}
