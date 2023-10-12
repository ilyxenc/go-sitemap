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
- **Delete**: удаляет URL в `urlsMap`.
- **End**: завершает создание карты сайта и записывает ее в указанный файл.

## Использование

Создание нового экземляра sitemap:

```go
sitemap := sm.NewSitemap()
```

Чтение файла sitemap.xml:

```go
sitemap, err := sm.Read("path/sitemap.xml")
if err != nil {
    // Обработка ошибки
}
```

Чтение данных ссылки:

```go
url, exists := sitemap.Get("https://example.com")
if !exists {
    // Действия при отсутствии
}
```

Обновление или добавление новой ссылки:

```go
sitemap.Upsert(sm.Url{Loc: "https://example.com"})
```

Удаление ссылки:
```go
ok := sitemap.Delete("https://example.com")
if !ok {
    // Действия при ошибке удаления
}
```

Запись данных в sitemap.xml:

```go
err := builder.End("path/sitemap.xml")
if err != nil {
    // Обработка ошибки
}
```

### Пример создания новой sitemap:

```go
package main

import (
	"fmt"
	"time"

	"github.com/ilyxenc/go-sitemap/sm"
)

func main() {
	sitemap := sm.NewSitemap()

	sitemap.Upsert(sm.Url{
		Loc:        "https://example.com",
		LastMod:    time.Now(),
		ChangeFreq: "daily",
		Priority:   1,
	})

	if err := sitemap.End("path/sitemap.xml"); err != nil {
		// Обработка ошибки
	}
}
```

### Пример чтения и редактирования sitemap:

```go
package main

import (
	"fmt"
	"time"

	"github.com/ilyxenc/go-sitemap/sm"
)

func main() {
	sitemap, err := sm.Read("path/sitemap.xml")
	if err != nil {
		// Обработка ошибки
	}

	sitemap.Upsert(sm.Url{
		Loc:        "https://example.com",
		LastMod:    time.Now(),
		ChangeFreq: "monthly",
		Priority:   0.8,
	})

	if err := sitemap.End("path/sitemap.xml"); err != nil {
		// Обработка ошибки
	}
}

```

## Тестирование

Для запуска тестов в корневой директории проекта выполните:

```bash
go test ./sm
```