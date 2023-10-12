package sm

import (
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"
)

// TestNewSitemap выполняет юнит-тестирование функции NewSitemap, которая создает новый
// экземпляр SitemapBuilder. Тест проверяет, что функция корректно создает и возвращает
// новый пустой экземпляр SitemapBuilder. Сценарий тестирования определен в структуре tests.
func TestNewSitemap(t *testing.T) {
	// Определение структуры для тестовых данных, включая поле want
	tests := []struct {
		name string
		want *SitemapBuilder
	}{
		{name: "Valid Sitemap Creation", want: &SitemapBuilder{urlsMap: make(map[string]Url)}},
	}

	// Выполнение цикла по всем сценариям тестирования
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызов функции NewSitemap и проверка возврата ожидаемого результата
			if got := NewSitemap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSitemap() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestRead выполняет юнит-тестирование функции Read, которая читает файл sitemap
// и возвращает экземпляр SitemapBuilder, заполненный данными из файла.
// Тест проверяет следующие сценарии:
// 1. Правильное чтение существующего файла
// 2. Обработка несуществующего файла.
// Каждый сценарий тестирования определен в структуре tests.
func TestRead(t *testing.T) {
	// Определение структуры для тестовых данных, включая поля fields и args
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *SitemapBuilder
		wantErr bool
	}{
		{
			name: "Valid sitemap file",
			args: args{filePath: "testdata/valid_sitemap.xml"},
			want: &SitemapBuilder{
				urlsMap: map[string]Url{
					"https://example.com": {
						Loc: "https://example.com",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Nonexistent file",
			args:    args{filePath: "testdata/nonexistent_sitemap.xml"},
			want:    nil,
			wantErr: true,
		},
	}

	// Выполнение цикла по всем сценариям тестирования
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызов функции Read и проверка возврата ожидаемого результата и ошибки (при наличии)
			got, err := Read(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSitemapBuilder_Get выполняет юнит-тестирование метода Get структуры SitemapBuilder.
// Тест проверяет, что метод правильно возвращает URL по его локации и корректно
// указывает наличие или отсутствие URL в URLs структуры SitemapBuilder.
// Каждый сценарий тестирования определен в структуре tests.
func TestSitemapBuilder_Get(t *testing.T) {
	// Определение структуры для тестовых данных, включая поля fields и args
	type fields struct {
		urlsMap map[string]Url
	}
	type args struct {
		loc string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Url
		exist  bool
	}{
		{
			name: "Existing URL",
			fields: fields{
				urlsMap: map[string]Url{
					"https://example.com": {
						Loc: "https://example.com",
					},
				},
			},
			args: args{
				loc: "https://example.com",
			},
			want: Url{
				Loc: "https://example.com",
			},
			exist: true,
		},
		{
			name: "Non-existing URL",
			fields: fields{
				urlsMap: make(map[string]Url),
			},
			args: args{
				loc: "https://example.com",
			},
			want:  Url{},
			exist: false,
		},
	}

	// Выполнение цикла по всем сценариям тестирования
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создание экземпляра SitemapBuilder с заданными данными
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}

			// Вызов метода Get и проверка, что URL и наличие/отсутствие URL корректно возвращаются
			got, exist := sb.Get(tt.args.loc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SitemapBuilder.Get() got = %v, want %v", got, tt.want)
			}
			if exist != tt.exist {
				t.Errorf("SitemapBuilder.Get() exist = %v, want %v", exist, tt.exist)
			}
		})
	}
}

// TestSitemapBuilder_Upsert выполняет юнит-тестирование метода Upsert структуры SitemapBuilder.
// Тест проверяет, что метод правильно вставляет новый URL в URLs структуры SitemapBuilder.
// Каждый сценарий тестирования определен в структуре tests.
func TestSitemapBuilder_Upsert(t *testing.T) {
	// Определение структуры для тестовых данных, включая поля fields и args
	type fields struct {
		urlsMap map[string]Url
	}
	type args struct {
		url Url
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]Url
	}{
		{
			name: "Insert new URL",
			fields: fields{
				urlsMap: make(map[string]Url),
			},
			args: args{
				url: Url{
					Loc: "https://example.com",
				},
			},
			want: map[string]Url{
				"https://example.com": {
					Loc: "https://example.com",
				},
			},
		},
	}

	// Выполнение цикла по всем сценариям тестирования
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создание экземпляра SitemapBuilder с заданными данными
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}

			// Вызов метода Upsert и проверка правильности вставки URL в URLs
			sb.Upsert(tt.args.url)

			if !reflect.DeepEqual(sb.urlsMap, tt.want) {
				t.Errorf("Post Upsert() urlsMap = %v, want %v", sb.urlsMap, tt.want)
			}
		})
	}
}

// TestSitemapBuilder_End выполняет юнит-тестирование метода End структуры SitemapBuilder.
// Тест проверяет, что метод правильно сохраняет Sitemap в XML-файл и обрабатывает следующие сценарии:
// 1. Правильные данные.
// 2. Некорректные пути к файлам.
// 3. Пустой набор URL
// Каждый сценарий тестирования определен в структуре tests.
func TestSitemapBuilder_End(t *testing.T) {
	// Создание временного файла для сохранения Sitemap и удаление его после завершения теста
	tmpFile, err := os.CreateTemp("./testdata", "sitemap-*.xml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Определение структуры для тестовых данных, включая поля fields и args
	type fields struct {
		urlsMap map[string]Url
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         bool ``
		expectedContent string
	}{
		{
			name: "Valid case with correct urlsMap",
			fields: fields{
				urlsMap: map[string]Url{
					"https://example.com": {
						Loc:        "https://example.com",
						ChangeFreq: "daily",
						Priority:   0.8,
					},
				},
			},
			args: args{
				filePath: tmpFile.Name(),
			},
			wantErr: false,
			expectedContent: `<?xml version="1.0" encoding="UTF-8"?>
				<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
					<url>
						<loc>https://example.com</loc>
						<lastmod>0001-01-01T00:00:00Z</lastmod>
						<changefreq>daily</changefreq>
						<priority>0.8</priority>
					</url>
				</urlset>`,
		},
		{
			name: "Invalid file path",
			fields: fields{
				urlsMap: map[string]Url{
					"https://example.com": {
						Loc:        "https://example.com",
						LastMod:    time.Now(),
						ChangeFreq: "daily",
						Priority:   0.8,
					},
				},
			},
			args: args{
				filePath: "./invalid-testdata/invalid_sitemap.xml",
			},
			wantErr:         true,
			expectedContent: "",
		},
		{
			name: "Empty urlsMap",
			fields: fields{
				urlsMap: map[string]Url{},
			},
			args: args{
				filePath: tmpFile.Name(),
			},
			wantErr: false,
			expectedContent: `<?xml version="1.0" encoding="UTF-8"?>
				<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"></urlset>`,
		},
	}

	// Выполнение цикла по всем сценариям тестирования
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создание экземпляра SitemapBuilder с заданными данными
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}

			// Вызов метода End и проверка на наличие ошибки, если она ожидается
			if err := sb.End(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("SitemapBuilder.End() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Проверка, что файл содержит ожидаемые данные, если ошибка не ожидается
			if !tt.wantErr {
				content, err := os.ReadFile(tt.args.filePath)
				if err != nil {
					t.Fatalf("Failed to read file: %v", err)
				}

				// Функция cleanWhitespace удаляет лишние пробелы и символы переноса строки
				cleanWhitespace := func(s string) string {
					whitespaceRegexp := regexp.MustCompile(`\s+`)
					return whitespaceRegexp.ReplaceAllString(s, "")
				}

				// Сравнение содержимого файла с ожидаемыми данными
				if cleanWhitespace(string(content)) != cleanWhitespace(tt.expectedContent) {
					t.Errorf("File content does not match expected. Got:\n%s\nWant:\n%s", content, tt.expectedContent)
				}
			}
		})
	}
}
