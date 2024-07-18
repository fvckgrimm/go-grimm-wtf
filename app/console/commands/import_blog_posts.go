package commands

import (
    "os"
    "path/filepath"
    "strings"
    "time"
    "bytes"

    "github.com/goravel/framework/facades"
    "goravel/app/models"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
    "gopkg.in/yaml.v2"
)

type ImportBlogPosts struct {
}

type frontMatter struct {
    Title       string    `yaml:"title"`
    Date        time.Time `yaml:"date"`
    Description string    `yaml:"description"`
    Tags        []string  `yaml:"tags"`
    Categories  []string  `yaml:"categories"`
    Draft       bool      `yaml:"draft"`
}

//Signature The name and signature of the console command.
func (receiver *ImportBlogPosts) Signature() string {
	return "import:posts"
}

//Description The console command description.
func (receiver *ImportBlogPosts) Description() string {
	return "Imports .md Post into Blog"
}

//Extend The console command extend.
func (receiver *ImportBlogPosts) Extend() command.Extend {
	return command.Extend{}
}

//Handle Execute the console command.
func (receiver *ImportBlogPosts) Handle(ctx console.Context) error {
    dir := "storage/app/blog-posts"
    files, err := os.ReadDir(dir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".md" {
            content, err := os.ReadFile(filepath.Join(dir, file.Name()))
            if err != nil {
                facades.Log().Error("Error reading file: " + err.Error())
                continue
            }

            parts := bytes.SplitN(content, []byte("---"), 3)
            if len(parts) != 3 {
                facades.Log().Error("Invalid file format: " + file.Name())
                continue
            }

            var fm frontMatter
            err = yaml.Unmarshal(parts[1], &fm)
            if err != nil {
                facades.Log().Error("Error parsing front matter: " + err.Error())
                continue
            }

            title := strings.TrimSuffix(file.Name(), ".md")
            slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))

            post := models.BlogPost{
                Title:       fm.Title,
                Slug:        slug,
                Content:     string(parts[2]),
                Date:        fm.Date,
                Description: fm.Description,
                Tags:        models.StringSlice(fm.Tags),
                Categories:  models.StringSlice(fm.Categories),
                Draft:       fm.Draft,
            }

            var existingPost models.BlogPost
            result := facades.Orm().Query().Where("slug", slug).First(&existingPost)
            
            if result.Error() != "" {
                // Post exists, update it
                existingPost.Title = post.Title
                existingPost.Content = post.Content
                existingPost.Date = post.Date
                existingPost.Description = post.Description
                existingPost.Tags = post.Tags
                existingPost.Categories = post.Categories
                existingPost.Draft = post.Draft
                result = facades.Orm().Query().Save(&existingPost)
            } else {
                // Post doesn't exist, create it
                result = facades.Orm().Query().Create(&post)
            }

            // Use Error directly without parentheses
            //result := facades.Orm().Query().Create(&post) 
            errStr := result.Error()
            if errStr != "" {
                facades.Log().Error("Error creating post: " + errStr)
            } else {
                status := "Imported"
                if fm.Draft {
                    status = "Imported as draft"
                }
                facades.Log().Info(status + " post: " + title)
            }
        }
    }
	
	return nil
}
