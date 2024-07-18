package controllers

import (
	"goravel/app/models"

    "path/filepath"
    "fmt"
    "os"
    "strings"
    "time"
    "bytes"
    "sort"
    "regexp"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

    "html/template"
    //"github.com/goravel/framework/facades"
    //"github.com/gomarkdown/markdown"
    //"github.com/gomarkdown/markdown/html"
    //"github.com/gomarkdown/markdown/parser"
    
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer/html"

    "gopkg.in/yaml.v2"
)

type BlogController struct {
	//Dependent services
}

type frontMatter struct {
    Title       string    `yaml:"title"`
    Date        time.Time `yaml:"date"`
    Description string    `yaml:"description"`
    Tags        []string  `yaml:"tags"`
    Categories  []string  `yaml:"categories"`
    Draft       bool      `yaml:"draft"`
}

func NewBlogController() *BlogController {
	return &BlogController{
		//Inject services
	}
}

func (r *BlogController) Index(ctx http.Context) http.Response {
	var posts []models.BlogPost 
    result := facades.Orm().Query().Where("draft", false).Order("created_at desc").Find(&posts)

    if len(posts) == 0 || result.Error() != "" {
        files, err := os.ReadDir("content/posts")
        if err == nil {
            for _, file := range files {
                if filepath.Ext(file.Name()) == ".md" {
                    slug := strings.TrimSuffix(file.Name(), ".md")
                    post, err := r.getPostFromFile(slug)
                    if err == nil && !post.Draft {
                        posts = append(posts, *post)
                    }
                }
            }
        }
    }

    sort.Slice(posts, func(i, j int) bool {
        return posts[i].Date.After(posts[j].Date)
    })

    template := "blog/index.tmpl"
    //if ctx.Value("isMobile").(bool) {
    //    template = "blog/mobile_index.tmpl"
    //}

    return ctx.Response().View().Make(template, map[string]interface{}{
        "posts": posts,
        "markdown": r.markdownFunc(),
    })
}	

func (c BlogController) Show(ctx http.Context) http.Response {
    slug := ctx.Request().Input("slug")
    var post *models.BlogPost
    var err error
    
    dbPost := models.BlogPost{}
    result := facades.Orm().Query().Where("slug", slug).First(&dbPost)
    if result.Error() == "" {
        if dbPost.Draft {
            return ctx.Response().Status(404).String("Post not found")
        }
        post = &dbPost
    } else {
        // If not found in the database, try to get it from the file
        post, err = c.getPostFromFile(slug)
        if err != nil || post.Draft {
            return ctx.Response().Status(404).String("Post not found")
        }
    }


    return ctx.Response().View().Make("blog/show.tmpl", map[string]interface{}{
        "post": post,
        "markdown": c.markdownFunc(),
    })
}

func (c *BlogController) markdownFunc() func(string) template.HTML {
    return func(input string) template.HTML {
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
            return template.HTML("")
        }

        // Add language class to code blocks
        re := regexp.MustCompile(`<pre><code class="language-(\w+)">`)
        result := re.ReplaceAllString(buf.String(), `<pre><code class="language-$1">`)

        return template.HTML(result)
    }
}

func (c BlogController) RecentPosts(ctx http.Context) http.Response {
    var posts []models.BlogPost
    limit := 5
    //models.DB().Order("created_at desc").Limit(3).Find(&posts)
    facades.Orm().Query()
    facades.Orm().Connection("sqlite").Query()
    facades.Log().Info("Attempting to fetch posts from database")
    result := facades.Orm().Query().Model(&models.BlogPost{}).Where("draft", false).Order("created_at desc").Limit(limit).Find(&posts)

    // If no posts found in the database or there was an error, read from files
    errStr := result.Error()
    if len(posts) == 0 || errStr != "" {
        facades.Log().Info("Fetching posts from files")
        posts = c.getRecentPostsFromFiles(limit)
    }

    if posts == nil {
        facades.Log().Info("No posts found, returning empty array")
        posts = []models.BlogPost{}
    }

    response := map[string]interface{}{
        "success": true,
        "response": posts,
    }

    //return ctx.Json(http.StatusOK, posts)
    return ctx.Response().Json(http.StatusOK, response)
}

func (c BlogController) getRecentPostsFromFiles(limit int) []models.BlogPost {
    var posts []models.BlogPost
    files, err := os.ReadDir("content/posts")
    if err != nil {
        return []models.BlogPost{}
    }

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".md" {
            slug := strings.TrimSuffix(file.Name(), ".md")
            post, err := c.getPostFromFile(slug)
            if err == nil && !post.Draft {
                posts = append(posts, *post)
            }
        }
    }

    // Sort posts by date (newest first)
    sort.Slice(posts, func(i, j int) bool {
        return posts[i].Date.After(posts[j].Date)
    })

    // Return only the most recent posts
    if len(posts) > limit {
        posts = posts[:limit]
    }

    return posts
}

func (c *BlogController) getPostFromFile(slug string) (*models.BlogPost, error) {
    filePath := filepath.Join("content", "posts", slug+".md")
    content, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    // parse front matter and content
    parts := bytes.SplitN(content, []byte("---"), 3)
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid file format for %s", slug)
    }

    // Parse front matter
    var fm frontMatter
    err = yaml.Unmarshal(parts[1], &fm)
    if err != nil {
        return nil, fmt.Errorf("error parsing front matter for %s: %v", slug, err)
    }

    // Create BlogPost instance
    post := &models.BlogPost{
        Title:       fm.Title,
        Slug:        slug,
        Content:     string(parts[2]),
        Date:        fm.Date,
        Description: fm.Description,
        Tags:        models.StringSlice(fm.Tags),
        Categories:  models.StringSlice(fm.Categories),
        Draft:       fm.Draft,
        CreatedAt:   fm.Date, // Assuming CreatedAt should be the same as the post date
    } 

    return post, nil
}
