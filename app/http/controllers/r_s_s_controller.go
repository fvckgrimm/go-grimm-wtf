package controllers

import (
    "goravel/app/models"
	"github.com/goravel/framework/contracts/http"
    "github.com/goravel/framework/facades"
    "github.com/gorilla/feeds"
    "time"

    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer/html"
    "bytes"
    "regexp"
)

type RSSController struct {
	//Dependent services
}

func NewRSSController() *RSSController {
	return &RSSController{
		//Inject services
	}
}

func (r *RSSController) Index(ctx http.Context) http.Response {
	var posts []models.BlogPost
    result := facades.Orm().Query().Where("draft", false).Order("created_at desc").Limit(20).Find(&posts)

    feed := &feeds.Feed{
        Title:       "Grimm's Graveyard",
        Link:        &feeds.Link{Href: "https://grimm.wtf"},
        Description: "Blog posts from Grimm's Graveyard",
        Author:      &feeds.Author{Name: "Grimm"},
        Created:     time.Now(),
    }

    if len(posts) == 0 || result.Error() != "" {
        // Fall back to reading from files
        blogController := NewBlogController()
        posts = blogController.getRecentPostsFromFiles(20)
    }

    for _, post := range posts {
        item := &feeds.Item{
            Title:       post.Title,
            Link:        &feeds.Link{Href: "https://grimm.wtf/blog/" + post.Slug},
            Description: post.Description,
            Author:      &feeds.Author{Name: "Grimm"},
            Created:     post.CreatedAt,
            Content:     r.markdownToHTML(post.Content),
        }
        feed.Items = append(feed.Items, item)
    }

    rss, err := feed.ToRss()
    if err != nil {
        return ctx.Response().String(500, "Error generating RSS feed")
    }

    return ctx.Response().Header("Content-Type", "application/rss+xml").String(200, rss)
}	


func (r *RSSController) markdownToHTML(input string) string {
    md := goldmark.New(
        goldmark.WithExtensions(
            extension.GFM,
            extension.Typographer,
            extension.Footnote,
        ),
        goldmark.WithParserOptions(
            parser.WithAutoHeadingID(),
        ),
        goldmark.WithRendererOptions(
            html.WithHardWraps(),
            html.WithXHTML(),
            html.WithUnsafe(),
        ),
    )

    var buf bytes.Buffer
    if err := md.Convert([]byte(input), &buf); err != nil {
        return ""
    }

    result := buf.String()

    // Clean up list rendering
    result = regexp.MustCompile(`(?m)^\s*<li>\s*`).ReplaceAllString(result, "<li>")
    result = regexp.MustCompile(`(?m)\s*</li>\s*$`).ReplaceAllString(result, "</li>")
    result = regexp.MustCompile(`(?m)^\s*<ul>\s*`).ReplaceAllString(result, "<ul>")
    result = regexp.MustCompile(`(?m)\s*</ul>\s*$`).ReplaceAllString(result, "</ul>")
    result = regexp.MustCompile(`(?m)^\s*<ol>\s*`).ReplaceAllString(result, "<ol>")
    result = regexp.MustCompile(`(?m)\s*</ol>\s*$`).ReplaceAllString(result, "</ol>")

    // Remove extra newlines
    result = regexp.MustCompile(`\n\s*\n`).ReplaceAllString(result, "\n")

    // Add language class to code blocks
    result = regexp.MustCompile(`<pre><code class="language-(\w+)">`).ReplaceAllString(result, `<pre><code class="language-$1">`)

    return result
}
