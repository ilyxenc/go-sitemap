package sm

import (
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestNewSitemap(t *testing.T) {
	tests := []struct {
		name string
		want *SitemapBuilder
	}{
		{name: "Valid Sitemap Creation", want: &SitemapBuilder{urlsMap: make(map[string]Url)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSitemap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSitemap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func TestSitemapBuilder_Get(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}
			got, exist := sb.Get(tt.args.loc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SitemapBuilder.Get() got = %v, want %v", got, tt.want)
			}
			if exist != tt.exist {
				t.Errorf("SitemapBuilder.Get() got1 = %v, want %v", exist, tt.exist)
			}
		})
	}
}

func TestSitemapBuilder_Upsert(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}
			sb.Upsert(tt.args.url)

			if !reflect.DeepEqual(sb.urlsMap, tt.want) {
				t.Errorf("Post Upsert() urlsMap = %v, want %v", sb.urlsMap, tt.want)
			}
		})
	}
}

func TestSitemapBuilder_End(t *testing.T) {
	tmpFile, err := os.CreateTemp("./testdata", "sitemap-*.xml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &SitemapBuilder{
				urlsMap: tt.fields.urlsMap,
			}
			if err := sb.End(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("SitemapBuilder.End() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				content, err := os.ReadFile(tt.args.filePath)
				if err != nil {
					t.Fatalf("Failed to read file: %v", err)
				}

				cleanWhitespace := func(s string) string {
					whitespaceRegexp := regexp.MustCompile(`\s+`)
					return whitespaceRegexp.ReplaceAllString(s, "")
				}

				if cleanWhitespace(string(content)) != cleanWhitespace(tt.expectedContent) {
					t.Errorf("File content does not match expected. Got:\n%s\nWant:\n%s", content, tt.expectedContent)
				}
			}
		})
	}
}
