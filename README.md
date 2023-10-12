# go-sitemap

Проект предоставляет инструменты для создания и редактирования карт сайта в формате XML.


## Основные компоненты

### SitemapBuilder

Основная структура для создания и обработки карты сайта.

#### Методы

- **NewSitemap**: создает новый экземпляр `SitemapBuilder`.
- **Read**: считывает файл карты сайта из указанного пути и возвращает экземпляр `SitemapBuilder`.
- **Get**: возвращает URL по указанному местоположению (loc).
- **Upsert**: обновляет или добавляет URL в `urlsMap`.
- **End**: завершает создание карты сайта и записывает ее в указанный файл.

## Использование

Для создания новой карты сайта:

```go
builder := sm.NewSitemap()

builder.Upsert(sm.Url{Loc: "http://example.com"})

err := builder.End("path/to/output.xml")
if err != nil {
    // Обработка ошибки
}
```

Для чтения и редактирования существующей карты сайта:

```go
builder, err := sm.Read("path/to/input.xml")
if err != nil {
    // Обработка ошибки
}

builder.Upsert(sm.Url{Loc: "http://example.com"})

url, exists := builder.Get("http://example.com")
if exists {
    // Действия при наличии
}

err = builder.End("path/to/output.xml")
if err != nil {
    // Обработка ошибки
}
```

## Тестирование

Для запуска тестов в корневой директории проекта выполните:

```bash
go test ./sm
```