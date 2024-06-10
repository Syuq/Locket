package frontend

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yourselfhosted/gomark/parser"
	"github.com/yourselfhosted/gomark/parser/tokenizer"
	"github.com/yourselfhosted/gomark/renderer"

	"github.com/Syuq/Locket/internal/util"
	"github.com/Syuq/Locket/server/profile"
	"github.com/Syuq/Locket/store"
)

const (
	// maxMetadataDescriptionLength is the maximum length of metadata description.
	maxMetadataDescriptionLength = 256
)

type FrontendService struct {
	Profile *profile.Profile
	Store   *store.Store
}

func NewFrontendService(profile *profile.Profile, store *store.Store) *FrontendService {
	return &FrontendService{
		Profile: profile,
		Store:   store,
	}
}

func (s *FrontendService) Serve(ctx context.Context, e *echo.Echo) {
	// Use echo static middleware to serve the built dist folder.
	// refer: https://github.com/labstack/echo/blob/master/middleware/static.go
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "dist",
		HTML5: true,
		Skipper: func(c echo.Context) bool {
			return util.HasPrefixes(c.Path(), "/api", "/lockets.api.v2", "/robots.txt", "/sitemap.xml", "/m/:name")
		},
	}))

	s.registerRoutes(e)
	s.registerFileRoutes(ctx, e)
}

func (s *FrontendService) registerRoutes(e *echo.Echo) {
	rawIndexHTML := getRawIndexHTML()

	e.GET("/m/:uid", func(c echo.Context) error {
		ctx := c.Request().Context()
		uid := c.Param("uid")
		locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
			UID: &uid,
		})
		if err != nil {
			return c.HTML(http.StatusOK, rawIndexHTML)
		}
		if locket == nil {
			return c.HTML(http.StatusOK, rawIndexHTML)
		}
		creator, err := s.Store.GetUser(ctx, &store.FindUser{
			ID: &locket.CreatorID,
		})
		if err != nil {
			return c.HTML(http.StatusOK, rawIndexHTML)
		}

		// Inject locket metadata into `index.html`.
		indexHTML := strings.ReplaceAll(rawIndexHTML, "<!-- lockets.metadata.head -->", generateLocketMetadata(locket, creator).String())
		indexHTML = strings.ReplaceAll(indexHTML, "<!-- lockets.metadata.body -->", fmt.Sprintf("<!-- lockets.locket.%d -->", locket.ID))
		return c.HTML(http.StatusOK, indexHTML)
	})
}

func (s *FrontendService) registerFileRoutes(ctx context.Context, e *echo.Echo) {
	workspaceGeneralSetting, err := s.Store.GetWorkspaceGeneralSetting(ctx)
	if err != nil {
		return
	}
	instanceURL := workspaceGeneralSetting.GetInstanceUrl()
	if instanceURL == "" {
		return
	}

	e.GET("/robots.txt", func(c echo.Context) error {
		robotsTxt := fmt.Sprintf(`User-agent: *
Allow: /
Host: %s
Sitemap: %s/sitemap.xml`, instanceURL, instanceURL)
		return c.String(http.StatusOK, robotsTxt)
	})

	e.GET("/sitemap.xml", func(c echo.Context) error {
		ctx := c.Request().Context()
		urlsets := []string{}
		// Append locket list.
		locketList, err := s.Store.ListLockets(ctx, &store.FindLocket{
			VisibilityList: []store.Visibility{store.Public},
		})
		if err != nil {
			return err
		}
		for _, locket := range locketList {
			urlsets = append(urlsets, fmt.Sprintf(`<url><loc>%s</loc></url>`, fmt.Sprintf("%s/m/%s", instanceURL, locket.UID)))
		}
		sitemap := fmt.Sprintf(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">%s</urlset>`, strings.Join(urlsets, "\n"))
		return c.XMLBlob(http.StatusOK, []byte(sitemap))
	})
}

func generateLocketMetadata(locket *store.Locket, creator *store.User) *Metadata {
	metadata := getDefaultMetadata()
	metadata.Title = fmt.Sprintf("%s(@%s) on Lockets", creator.Nickname, creator.Username)
	if locket.Visibility == store.Public {
		tokens := tokenizer.Tokenize(locket.Content)
		nodes, _ := parser.Parse(tokens)
		description := renderer.NewStringRenderer().Render(nodes)
		if len(description) == 0 {
			description = locket.Content
		}
		if len(description) > maxMetadataDescriptionLength {
			description = description[:maxMetadataDescriptionLength] + "..."
		}
		metadata.Description = description
	}

	return metadata
}

func getRawIndexHTML() string {
	bytes, _ := os.ReadFile("dist/index.html")
	return string(bytes)
}

type Metadata struct {
	Title       string
	Description string
	ImageURL    string
}

func getDefaultMetadata() *Metadata {
	return &Metadata{
		Title:       "Lockets",
		Description: "LOCKET IMAGE SHARING APP.",
		ImageURL:    "/logo.webp",
	}
}

func (m *Metadata) String() string {
	metadataList := []string{
		fmt.Sprintf(`<meta name="description" content="%s" />`, m.Description),
		fmt.Sprintf(`<meta property="og:title" content="%s" />`, m.Title),
		fmt.Sprintf(`<meta property="og:description" content="%s" />`, m.Description),
		fmt.Sprintf(`<meta property="og:image" content="%s" />`, m.ImageURL),
		`<meta property="og:type" content="website" />`,
		// Twitter related fields.
		fmt.Sprintf(`<meta property="twitter:title" content="%s" />`, m.Title),
		fmt.Sprintf(`<meta property="twitter:description" content="%s" />`, m.Description),
		fmt.Sprintf(`<meta property="twitter:image" content="%s" />`, m.ImageURL),
		`<meta name="twitter:card" content="summary" />`,
		`<meta name="twitter:creator" content="lockets" />`,
	}
	return strings.Join(metadataList, "\n")
}
